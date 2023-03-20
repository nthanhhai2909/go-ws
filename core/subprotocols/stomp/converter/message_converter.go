package converter

import (
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[P interface{}] interface {
	FromMessage(msg smsg.IMessage, target P) error
	ToMessage(payload P, headers header.Header) (error, smsg.IMessage)
}
