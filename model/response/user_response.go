package response

type (
	// UserResponse User Request
	// @Description User information server responses
	UserResponse struct {
		ID        uint   `json:"id" validate:"required"`
		Username  string `json:"login" validate:"required"`
		AvatarUrl string `json:"avatar_url" validate:"url"`
		Email     string `json:"email" validate:"email"`
		Bio       string `json:"bio"`
	}
)
