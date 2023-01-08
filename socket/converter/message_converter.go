package converter

import (
	"mem-ws/socket/header"
	"mem-ws/socket/stomp/stompmsg"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}] interface {
	FromMessage(msg stompmsg.Message[P], target P) error
	ToMessage(payload P, headers header.Headers) (error, stompmsg.Message[P])
}

func NewStringMessageConverter[P interface{}]() MessageConverter[P] {
	return &stringMessageConverter[P]{}
}
