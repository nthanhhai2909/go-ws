package smsg

import (
	"mem-ws/socket/core/header"
)

type IMessage[P interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *header.Headers
}
