package web

type SignInRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiredAt   int64  `json:"expired_at"`
}

type ProfileResponse struct {
	PhotoURL string `json:"photo_url"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Bio      string `json:"bio"`
}
