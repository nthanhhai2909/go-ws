package channel

import "mem-ws/socket/stomp/msg"

type Channel interface {
	Subscribe(handler msg.Handler) error
	Unsubscribe(handler msg.Handler) error
	Send(message msg.Message[interface{}]) error
}
