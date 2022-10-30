package msg

import (
	"mem-ws/socket/header"
)

type Message[P interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *header.Headers
}
