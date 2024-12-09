package web

type UpdateProfileRequest struct {
	Gender   string `json:"gender" validate:"omitempty,oneof=MALE FEMALE"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	PhotoURL string `json:"photo_url" validate:"omitempty,url"`
}
