package response

type (
	// TagResponse Tag Request
	// @Description Tag information server responses
	TagResponse struct {
		ID   uint   `json:"id" validate:"required,numeric"`
		Name string `json:"name" validate:"required"`
	}
)
