package outbound

import (
	"mem-ws/core/message"
)

type OutboundChannel[T interface{}] interface {
	Send(message message.Message[T], timeout int64) error
}
