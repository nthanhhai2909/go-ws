package converter

import (
	"mem-ws/core/message"
)

// MessageConverter TODO HGA: UPDATING
type MessageConverter[T interface{}] interface {
	FromMessage()
	ToMessage() message.Message[T]
}
