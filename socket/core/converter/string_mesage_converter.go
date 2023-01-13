package converter

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/smsg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg smsg.Message[P], target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, smsg.Message[P]) {
	return nil, nil
}
