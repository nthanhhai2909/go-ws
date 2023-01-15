package channel

import "mem-ws/socket/core/stomp/smsg"

type Channel interface {
	Subscribe(handler smsg.Handler) error
	Unsubscribe(handler smsg.Handler) error
	Send(message smsg.IMessage) error
}
