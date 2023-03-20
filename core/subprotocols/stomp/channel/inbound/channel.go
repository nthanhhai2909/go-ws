package inbound

import (
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

type IChannel interface {
	Subscribe(msg smsg.IMessage, session session.ISession) error
	Unsubscribe(msg smsg.IMessage, session session.ISession) error
	Send(message smsg.IMessage) error
}
