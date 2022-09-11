package dto

type WSRequest struct {
	ID      string `json:"id"`
	Action  Action `json:"action"`
	Payload string `json:"payload"`
}

type UserBroadCast struct {
	Destination string `json:"destination"`
	Data        string `json:"data"`
}
