package message

type Message[P interface{}, H interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *Headers[H]
}
