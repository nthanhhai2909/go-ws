package converter

import (
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
)

type stringMessageConverter[P interface{}] struct {
}

func (converter *stringMessageConverter[P]) FromMessage(msg smsg.IMessage, target P) error {
	return nil
}

func (converter *stringMessageConverter[P]) ToMessage(payload P, headers header.Header) (error, smsg.IMessage) {
	return nil, nil
}
