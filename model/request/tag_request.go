package request

type (
	// TagRequest Tag Request
	// @Description Necessary tag information when request
	TagRequest struct {
		Name string `json:"name" validate:"required"`
	}
)
