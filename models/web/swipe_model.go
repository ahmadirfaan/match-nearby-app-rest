package web

type SwipeRequest struct {
	ProfileId string `json:"profile_id"  validate:"required,min=24,max=24"`
	Action    bool   `json:"action"  validate:"required"`
}
