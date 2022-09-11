package dto

type WSRequestDTO struct {
	ID      string `json:"id"`
	Action  Action `json:"action"`
	Payload string `json:"payload"`
}
