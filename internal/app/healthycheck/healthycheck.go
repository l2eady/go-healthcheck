package healthycheck

import (
	"crypto/tls"
	"encoding/csv"
	"go-healthcheck/internal/app"
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
	Conf   *app.Configs
}

// NewHealthyCheckService will create a healthy check service layer
func NewHealthyCheckService(maxTimeOut time.Duration, conf *app.Configs) HealthyCheckService {
	return &healthyCheckServiceImpl{
		Caller: &lhttp.Caller{
			Body:   nil,
			Header: map[string]string{},
			Client: &http.Client{Timeout: maxTimeOut, Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}},
		},
		Conf: conf,
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
	pool.Result.EndAt = time.Now()
	service.SendReport(pool.Result)

	return pool.Result
}
func (service *healthyCheckServiceImpl) SendReport(report *models.HealthyCheckReport) error {
	type reportRequest struct {
		TotalWebsites         int   `json:"total_websites"`
		TotalSuccess          int   `json:"success"`
		TotalFailure          int   `json:"failure"`
		TotalTimeInNanoSecond int64 `json:"total_time"`
	}
	req := reportRequest{
		TotalWebsites:         len(report.Data),
		TotalSuccess:          report.TotalSuccess,
		TotalFailure:          report.TotalFailure,
		TotalTimeInNanoSecond: report.TotalTimeUsedInNano(),
	}

	url := service.Conf.ReportService.Address + service.Conf.ReportService.ReportEndPoint
	service.Caller.SetURL(url)
	service.Caller.SetBody(req)
	service.Caller.SetHeader(map[string]string{"Authorization": service.Conf.ReportService.AccessToken})
	_, err := service.Caller.POST()
	if err != nil {
		log.Printf("the system can't send the report, err : %v\n", err)
	}
	return err
}
