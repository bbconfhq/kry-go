package request

type (
	// ProblemRequest Problem Request
	// @Description Necessary problem information when request
	ProblemRequest struct {
		Title       string  `json:"title" validate:"required"`
		Content     string  `json:"content" validate:"required"`
		Note        string  `json:"note"`
		TimeLimit   float64 `json:"time_limit" validate:"required;numeric"`
		MemoryLimit uint    `json:"memory_limit" validate:"required;numeric"`
	}
)
