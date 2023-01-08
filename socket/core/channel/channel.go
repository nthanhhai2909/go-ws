package channel

import (
	stompmsg2 "mem-ws/socket/core/stomp/stompmsg"
)

type Channel interface {
	Subscribe(handler stompmsg2.Handler) error
	Unsubscribe(handler stompmsg2.Handler) error
	Send(message stompmsg2.Message[interface{}]) error
}
