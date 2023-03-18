package converter

import (
	"mem-ws/core/header"
	"mem-ws/core/subprotocols/stomp/smsg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg smsg.IMessage, target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Headers) (error, smsg.IMessage) {
	return nil, nil
}
