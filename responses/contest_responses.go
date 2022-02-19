package responses

import "time"

type (
	ContestResponse struct {
		ID         uint
		Title      string
		Content    string
		ProblemIds []uint
		StartedAt  time.Time
		EndedAt    time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
