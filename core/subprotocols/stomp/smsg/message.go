package smsg

import (
	"mem-ws/core/header"
)

type IMessage interface {
	GetPayload() []byte
	GetMessageHeaders() *header.Headers
}
