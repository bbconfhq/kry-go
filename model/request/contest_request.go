package request

import "time"

type (
	// ContestRequest Contest Request
	// @Description Necessary contest information when request
	ContestRequest struct {
		Title      string    `json:"title" validate:"required"`
		Content    string    `json:"content" validate:"required,max=191"`
		ProblemIds []uint    `json:"problem_ids" validate:"required"`
		StartedAt  time.Time `json:"started_at" validate:"required"`
		EndedAt    time.Time `json:"ended_at" validate:"required"`
	}
)
