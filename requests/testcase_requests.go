package requests

type (
	TestcaseRequest struct {
		Visible   bool   `json:"visible" validate:"required,boolean"`
		Input     string `json:"input" validate:"required"`
		Output    string `json:"output" validate:"required"`
		ProblemID uint   `json:"problem_id" validate:"required"`
	}
)
