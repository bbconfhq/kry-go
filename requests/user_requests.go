package requests

type (
	UserRequest struct {
		ID       uint   `json:"id" validate:"required"`
		Username string `json:"login" validate:"required"`
		Email    string `json:"email" validate:"email"`
		Bio      string `json:"bio"`
	}
)
