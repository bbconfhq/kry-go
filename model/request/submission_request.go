package request

type (
	// SubmissionRequest Submission Request
	// @Description Necessary submission information when request
	SubmissionRequest struct {
		Result    uint   `json:"result" validate:"required,numeric"`
		Language  string `json:"language" validate:"required"`
		Code      string `json:"code" validate:"required"`
		UserID    uint   `json:"user_id" validate:"required,numeric"`
		ProblemID uint   `json:"problem_id" validate:"required,numeric"`
	}
)
