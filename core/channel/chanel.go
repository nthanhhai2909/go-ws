package channel

import "mem-ws/core"

type Channel[T interface{}] interface {
	Send(message core.Message[T], timeout int64) error
}
