package smsg

import (
	"mem-ws/socket/core/header"
)

type IMessage interface {
	GetPayload() []byte
	GetMessageHeaders() *header.Headers
}
