package models

import "time"

type HealthyCheckRequest struct {
	URL string
}

type HealthyCheckResponse struct {
	URL     string
	Status  bool
	StartAt time.Time
	EndAt   time.Time
}
