package outbound

import (
	"mem-ws/core/stomp/msg"
)

type Channel[P interface{}] interface {
	Send(message msg.Message[P], timeout int64) error
}
