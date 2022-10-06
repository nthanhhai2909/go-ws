package converter

import "mem-ws/core"

// MessageConverter TODO HGA: UPDATING
type MessageConverter[T interface{}] interface {
	FromMessage()
	ToMessage() core.Message[T]
}
