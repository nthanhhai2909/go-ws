package smsg

import (
	"mem-ws/core/subprotocols/stomp/header"
)

type GenericMessage struct {
	Payload []byte
	Headers *header.Header
}

func (msg *GenericMessage) GetPayload() []byte {
	return msg.Payload
}

func (msg *GenericMessage) GetMessageHeaders() *header.Header {
	return msg.Headers
}

func NewGenericMessage(headerProperties map[string]string, payload []byte) IMessage {
	return &GenericMessage{Headers: header.WithProperties(headerProperties), Payload: payload}
}
