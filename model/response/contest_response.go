package response

import "time"

type (
	// ContestResponse Contest Response
	// @Description Contest information server responses
	ContestResponse struct {
		ID         uint
		Title      string
		Content    string
		ProblemIds []uint
		StartedAt  time.Time
		EndedAt    time.Time
	}
)
