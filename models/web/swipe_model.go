package web

type SwipeRequest struct {
	UserID string `json:"user_id"  validate:"required,min=24,max=30"`
	Action bool   `json:"action"`
}

type GetProfileResponse struct {
	Data           []ProfileModelResponse `json:"data"`
	RemainingQuota uint16                 `json:"remaining_quota"`
}

type ProfileModelResponse struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Photo  string `json:"photo"`
	Bio    string `json:"bio"`
}
