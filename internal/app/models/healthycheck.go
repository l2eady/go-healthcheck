package models

import (
	"time"
)

type HealthyCheckRequest struct {
	URL string
}

type HealthyCheckResponse struct {
	URL       string
	IsSuccess bool
	StartAt   time.Time
	EndAt     time.Time
}

type HealthyCheckReport struct {
	StartAt      time.Time
	Data         []HealthyCheckResponse
	TotalSuccess int
	TotalFailure int
	EndAt        time.Time
}

func (h HealthyCheckReport) TotalTimeUsedInNano() int64 {
	return h.EndAt.Sub(h.StartAt).Nanoseconds()
}
