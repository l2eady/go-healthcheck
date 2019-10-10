package healthycheck

import (
	"fmt"
	"go-healthcheck/internal/app/models"
	"sync"
)

type Worker struct {
	ID           string
	ChanelJob    chan models.HealthyCheckRequest
	ChanelResult chan models.HealthyCheckResponse
}

func (w *Worker) Do(service HealthyCheckService, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range w.ChanelJob {
		fmt.Printf("%s started working on job: %s\n", w.ID, job.URL)
		resp, _ := service.Ping(job)
		fmt.Printf("%s finish working on job: %s\n", w.ID, job.URL)
		w.ChanelResult <- resp
	}
	fmt.Println(w.ID, "don't have anyjob, then go home")
}
