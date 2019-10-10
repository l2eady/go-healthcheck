package healthycheck

import (
	"go-healthcheck/internal/app/lhttp/mocks"
	"go-healthcheck/internal/app/models"
	"sync"
	"testing"
)

func TestWorkerExpectDoJobReturnResultInChannel(t *testing.T) {
	mockCaller := &mocks.MockLHTTPCaller{MockGETReturnSuccess: nil}
	mockService := &healthyCheckServiceImpl{Caller: mockCaller}
	chResult := make(chan models.HealthyCheckResponse)
	chJob := make(chan models.HealthyCheckRequest)
	worker := &Worker{
		ID:           "worker_test",
		ChanelJob:    chJob,
		ChanelResult: chResult,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go worker.Do(mockService, wg)
	chJob <- models.HealthyCheckRequest{URL: "https://linecorp.com"}
	close(chJob)
	result := <-chResult
	if !result.IsSuccess {
		t.Errorf("Expecte true, but got %v", result.IsSuccess)
	}
	wg.Wait()
}
