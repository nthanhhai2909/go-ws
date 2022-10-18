package converter

import (
	"mem-ws/core/stomp/header"
	message2 "mem-ws/core/stomp/msg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg message2.Message[P], target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, message2.Message[P]) {
	return nil, nil
}
