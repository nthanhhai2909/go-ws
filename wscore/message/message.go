package message

type Message[T interface{}] interface {
	GetPayload() T
	GetMessageHeaders() Headers[string, interface{}]
}
