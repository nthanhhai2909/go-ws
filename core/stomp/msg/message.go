package msg

import "mem-ws/core/stomp/header"

type Message[P interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *header.Headers
}
