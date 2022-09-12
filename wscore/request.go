package wscore

type WSRequest struct {
	ID      string `json:"id"`
	Action  Action `json:"action"`
	Payload string `json:"payload"`
}
