package smsg

import (
	"mem-ws/socket/core/header"
)

type Message[P interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *header.Headers
}
