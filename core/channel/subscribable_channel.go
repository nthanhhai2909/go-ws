package channel

import (
	"mem-ws/core"
	"mem-ws/core/wserror"
)

type SubscribableChannel[T interface{}] struct {
	MessageHandler core.Subscriber[T]
}

// Send TODO HGA: CONSIDER TIMEOUT CONSTRAINTS
func (chann *SubscribableChannel[T]) Send(message core.Message[T], timeout int64) error {
	if message.GetPayload() == nil {
		return wserror.IllegalArgument{Message: "Message payload must not be null"}
	}
	return nil
}
