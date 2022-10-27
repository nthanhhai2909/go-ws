package socket

type WebsocketMessage[T interface{}] interface {
	GetPayload() T
	GetPayloadLength() int
	// TODO HGA
	IsLast()
}
