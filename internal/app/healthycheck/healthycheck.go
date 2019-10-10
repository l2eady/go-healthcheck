package healthycheck

import (
	"encoding/csv"
	"go-healthcheck/internal/app/lhttp"
	"go-healthcheck/internal/app/models"
	"io"
	"log"
	"net/http"
	"time"
)

type HealthyCheckService interface {
	Ping(models.HealthyCheckRequest) (models.HealthyCheckResponse, error)
	HealthyCheckEndPointFromCSVFile(io.Reader, int) *models.HealthyCheckReport
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
		URL:       req.URL,
		StartAt:   time.Now(),
		IsSuccess: false,
	}

	service.Caller.SetURL(req.URL)
	_, err := service.Caller.GET()

	// after ping
	resp.EndAt = time.Now()
	if err != nil {
		return resp, err
	}
	resp.IsSuccess = true
	return resp, nil
}

func (service *healthyCheckServiceImpl) HealthyCheckEndPointFromCSVFile(reader io.Reader, maxWorker int) *models.HealthyCheckReport {
	csvReader := csv.NewReader(reader)
	pool := service.NewPool(maxWorker)
	pool.Start()
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Read file error: %#v \n", err)
			continue
		}
		pool.ChannelJob <- models.HealthyCheckRequest{URL: record[0]}
	}
	close(pool.ChannelJob)
	<-pool.Done
	return pool.Result
}
