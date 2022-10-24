package inbound

import (
	"mem-ws/core/stomp/msg"
)

// Channel TODO HGA WILL ADAPT
type Channel[P interface{}] interface {
	Send(message msg.Message[P]) error
}
