package channel

import (
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

type Inbound interface {
	Subscribe(session session.ISession) error
	Unsubscribe(session session.ISession) error
	Send(message smsg.IMessage) error
}
