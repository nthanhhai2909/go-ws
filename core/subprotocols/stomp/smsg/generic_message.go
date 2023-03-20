package smsg

import (
	"mem-ws/core/subprotocols/stomp/header"
)

type GenericMessage[P interface{}] struct {
	Payload P
	Headers *header.Header
}

func (msg *GenericMessage[P]) GetPayload() P {
	return msg.Payload
}

func (msg *GenericMessage[P]) GetMessageHeaders() *header.Header {
	return msg.Headers
}
