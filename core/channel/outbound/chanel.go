package outbound

import (
	"mem-ws/core/message"
)

type Channel[P interface{}, H interface{}] interface {
	Send(message message.Message[P, H], timeout int64) error
}
