package converter

import (
	"mem-ws/core/stomp/header"
	message2 "mem-ws/core/stomp/msg"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}] interface {
	FromMessage(msg message2.Message[P], target P) error
	ToMessage(payload P, headers header.Headers) (error, message2.Message[P])
}

func NewStringMessageConverter[P interface{}]() MessageConverter[P] {
	return &stringMessageConverter[P]{}
}
