package channel

import "mem-ws/socket/stomp/stompmsg"

type Channel interface {
	Subscribe(handler stompmsg.Handler) error
	Unsubscribe(handler stompmsg.Handler) error
	Send(message stompmsg.Message[interface{}]) error
}
