package web

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Gender   string `json:"gender" validate:"required,oneof=MALE FEMALE"`
	Name     string `json:"name" validate:"required"`
}
