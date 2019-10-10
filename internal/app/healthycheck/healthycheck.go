package healthycheck

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/internal/app/lhttp"
	"go-healthcheck/internal/app/models"
	"io"
	"log"
	"net/http"
	"time"
)

type HealthyCheckService interface {
	Ping(req models.HealthyCheckRequest) (models.HealthyCheckResponse, error)
	HealthyCheckEndPointFromCSVFile(io.Reader) ([]models.HealthyCheckResponse, error)
}
type healthyCheckServiceImpl struct {
	Caller lhttp.HttpCaller
}

// NewHealthyCheckService will create a healthy check service layer
func NewHealthyCheckService(maxTimeOut time.Duration) HealthyCheckService {
	return &healthyCheckServiceImpl{
		Caller: &lhttp.Caller{
			Body:   nil,
			Header: map[string]string{},
			Client: &http.Client{Timeout: maxTimeOut},
		},
	}
}

func (service healthyCheckServiceImpl) Ping(req models.HealthyCheckRequest) (models.HealthyCheckResponse, error) {
	resp := models.HealthyCheckResponse{
		URL:     req.URL,
		StartAt: time.Now(),
	}

	service.Caller.SetURL(req.URL)
	_, err := service.Caller.GET()

	// after ping
	resp.EndAt = time.Now()
	if err != nil {
		fmt.Printf("Ping to %s failed\n", resp.URL)
		resp.Status = false
		return resp, err
	}
	resp.Status = true
	fmt.Printf("Ping to %s successfuly\n", resp.URL)
	return resp, nil
}

func (service *healthyCheckServiceImpl) HealthyCheckEndPointFromCSVFile(reader io.Reader) ([]models.HealthyCheckResponse, error) {
	csvReader := csv.NewReader(reader)
	pool := service.NewPool(100)
	pool.Start()
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			// switch err {
			// case io.EOF:
			// 	break
			// default:
			log.Printf("Read file error: %#v \n", err)
			return nil, err
			// }
		}
		fmt.Println("WTF")
		pool.ChannelJob <- models.HealthyCheckRequest{URL: record[0]}
	}
	close(pool.ChannelJob)
	<-pool.Done
	return pool.Results, nil
}
