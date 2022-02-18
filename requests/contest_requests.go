package requests

import "time"

type (
	ContestRequest struct {
		Title     string    `json:"title" validate:"required"`
		Content   string    `json:"content" validate:"required,max=191"`
		Problems  []uint    `json:"problems" validate:"required"`
		StartedAt time.Time `json:"started_at" validate:"required"`
		EndedAt   time.Time `json:"ended_at" validate:"required"`
	}
)
