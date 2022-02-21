package response

import (
	"time"
)

type (
	// ProblemResponse Problem Response
	// @Description Problem information server responses
	ProblemResponse struct {
		ID          uint
		Title       string
		Content     string
		Note        string
		TimeLimit   float64
		MemoryLimit uint
		SubmitCount uint
		AcceptCount uint
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Tags        []string
	}
)
