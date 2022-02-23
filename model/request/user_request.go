package request

type (
	// UserRequest User Request
	// @Description Necessary user information when request
	UserRequest struct {
		ID        uint   `json:"id" validate:"required"`
		Username  string `json:"login" validate:"required"`
		AvatarUrl string `json:"avatar_url" validate:"url"`
		Email     string `json:"email" validate:"email"`
		Bio       string `json:"bio"`
	}
)
