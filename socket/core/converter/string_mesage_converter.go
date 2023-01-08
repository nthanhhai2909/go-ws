package converter

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/stompmsg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg stompmsg.Message[P], target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, stompmsg.Message[P]) {
	return nil, nil
}
