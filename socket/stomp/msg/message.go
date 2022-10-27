package msg

import "mem-ws/socket/stomp/header"

type Message[P interface{}] interface {
	GetPayload() P
	GetMessageHeaders() *header.Headers
}
