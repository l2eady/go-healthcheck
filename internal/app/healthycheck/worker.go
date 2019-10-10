package healthycheck

import (
	"go-healthcheck/internal/app/models"
	"sync"
)

type Worker struct {
	ID           string
	ChanelJob    chan models.HealthyCheckRequest
	ChanelResult chan models.HealthyCheckResponse
}

// Do will call ping function and send response back to channel
func (w *Worker) Do(service HealthyCheckService, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range w.ChanelJob {
		resp, _ := service.Ping(job)
		w.ChanelResult <- resp
	}
}
