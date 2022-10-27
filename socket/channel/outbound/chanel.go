package outbound

import (
	"mem-ws/socket/stomp/msg"
)

type Channel[P interface{}] interface {
	Send(message msg.Message[P], timeout int64) error
}
