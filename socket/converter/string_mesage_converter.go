package converter

import (
	"mem-ws/socket/header"
	"mem-ws/socket/stomp/msg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg msg.Message[P], target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, msg.Message[P]) {
	return nil, nil
}
