package outbound

import (
	"mem-ws/core/message"
)

type Channel[T interface{}] interface {
	Send(message message.Message[T], timeout int64) error
}
