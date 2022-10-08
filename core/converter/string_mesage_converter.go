package converter

import (
	"mem-ws/core/message"
)

type stringMessageConverter[P interface{}, T interface{}] struct {
}

func (converter *stringMessageConverter[P, T]) FromMessage(msg message.Message[P, T], target T) error {
	return nil
}

func (converter *stringMessageConverter[P, T]) ToMessage(payload P, headers message.Headers[T]) (error, message.Message[P, T]) {
	return nil, nil
}
