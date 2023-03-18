package message

type IMessage interface {
	GetPayload() []byte
	GetPayloadLength() int
	IsLast() bool
}
