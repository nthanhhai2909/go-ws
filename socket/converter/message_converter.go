package converter

import (
	"mem-ws/socket/header"
	"mem-ws/socket/stomp/msg"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}] interface {
	FromMessage(msg msg.Message[P], target P) error
	ToMessage(payload P, headers header.Headers) (error, msg.Message[P])
}

func NewStringMessageConverter[P interface{}]() MessageConverter[P] {
	return &stringMessageConverter[P]{}
}
