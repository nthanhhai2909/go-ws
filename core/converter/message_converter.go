package converter

import (
	"mem-ws/core/message"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}, H interface{}] interface {
	FromMessage(msg message.Message[P, H], target H) error
	ToMessage(payload P, headers message.Headers[H]) (error, message.Message[P, H])
}

func NewStringMessageConverter[P interface{}, T interface{}]() MessageConverter[P, T] {
	return &stringMessageConverter[P, T]{}
}
