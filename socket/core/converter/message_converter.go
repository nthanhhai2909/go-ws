package converter

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/smsg"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}] interface {
	FromMessage(msg smsg.IMessage[P], target P) error
	ToMessage(payload P, headers header.Headers) (error, smsg.IMessage[P])
}

func NewStringMessageConverter[P interface{}]() MessageConverter[P] {
	return &stringMessageConverter[P]{}
}
