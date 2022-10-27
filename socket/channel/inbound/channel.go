package inbound

import (
	"mem-ws/socket/stomp/msg"
)

type Channel[P interface{}] interface {
	Subscribe(handler msg.Handler[P]) error
	Unsubscribe(handler msg.Handler[P]) error
	Send(message msg.Message[P]) error
}
