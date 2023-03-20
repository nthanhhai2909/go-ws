package smsg

import (
	"mem-ws/core/subprotocols/stomp/header"
)

type IMessage interface {
	GetPayload() []byte
	GetMessageHeaders() *header.Header
}
