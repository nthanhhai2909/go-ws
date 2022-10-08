package message

type GenericMessage[P interface{}, H interface{}] struct {
	Payload P
	Headers *Headers[H]
}

func (msg *GenericMessage[P, T]) GetPayload() P {
	return msg.Payload
}

func (msg *GenericMessage[P, T]) GetMessageHeaders() *Headers[T] {
	return msg.Headers
}
