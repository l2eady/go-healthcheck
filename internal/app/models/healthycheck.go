package models

import (
	"time"
)

type HealthyCheckRequest struct {
	URL string
}

type HealthyCheckResponse struct {
	URL     string
	IsSuccess  bool
	StartAt time.Time
	EndAt   time.Time
}

type HealthyCheckReport struct {
	Data         []HealthyCheckResponse
	CountSuccess int
	CountFailure int
}
