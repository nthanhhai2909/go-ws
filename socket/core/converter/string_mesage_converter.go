package converter

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/smsg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg smsg.IMessage[P], target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, smsg.IMessage[P]) {
	return nil, nil
}
