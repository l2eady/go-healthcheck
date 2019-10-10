package mocks

import "go-healthcheck/internal/app/models"

type MockHealthyCheckService struct {
	MockHealthyCheckResponse    models.HealthyCheckResponse
	MockHealthyCheckResponseErr error
}

func (m *MockHealthyCheckService) Ping(url string) (ok models.HealthyCheckResponse, err error) {
	return m.MockHealthyCheckResponse, m.MockHealthyCheckResponseErr
}
