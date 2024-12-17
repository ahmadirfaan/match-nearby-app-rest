package web

type SwipeRequest struct {
	UserID string `json:"user_id"  validate:"required,min=24,max=24"`
	Action bool   `json:"action"  validate:"required"`
}
