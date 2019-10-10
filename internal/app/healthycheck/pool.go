package healthycheck

import (
	"fmt"
	"go-healthcheck/internal/app/models"
	"sync"
)

type pool struct {
	Done          chan bool
	ChannelJob    chan models.HealthyCheckRequest
	ChanelResult  chan models.HealthyCheckResponse
	maxGoRoutines int
	Results       []models.HealthyCheckResponse
	HealthyCheckService
}

func (h *healthyCheckServiceImpl) NewPool(maxGoRoutines int) *pool {
	return &pool{
		Done:                make(chan bool),
		ChannelJob:          make(chan models.HealthyCheckRequest, maxGoRoutines),
		ChanelResult:        make(chan models.HealthyCheckResponse, maxGoRoutines),
		maxGoRoutines:       maxGoRoutines,
		Results:             make([]models.HealthyCheckResponse, maxGoRoutines),
		HealthyCheckService: h,
	}
}

// Start will spawning new workers
func (p *pool) Start() {
	go func() {
		defer close(p.ChanelResult)
		wg := &sync.WaitGroup{}
		fmt.Printf("Worker Pool spawning new goRoutines, total: [%d]\n", p.maxGoRoutines)
		for i := 1; i <= p.maxGoRoutines; i++ {
			wg.Add(1)
			worker := &Worker{
				ID:           fmt.Sprintf("[Worker ID]: %d", i),
				ChanelJob:    p.ChannelJob,
				ChanelResult: p.ChanelResult,
			}
			// create new worker in new gorutine
			go worker.Do(p.HealthyCheckService, wg)
		}
		fmt.Println("Worker Pool done spawning work goRoutines")
		wg.Wait()
	}()
	p.collect()
}

func (p *pool) collect() {
	go func() {
		defer func() { p.Done <- true }()
		for result := range p.ChanelResult {
			p.Results = append(p.Results, result)
		}
	}()
}
