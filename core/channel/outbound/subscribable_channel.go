package outbound

import (
	"mem-ws/core/message"
	"mem-ws/core/wserror"
)

type SubscribableChannel[T interface{}] struct {
	MessageHandler []message.Handler[T]
}

func NewSubscribableChannel(handler []message.Handler[interface{}]) *SubscribableChannel[interface{}] {
	return &SubscribableChannel[interface{}]{MessageHandler: handler}
}

func NewSingleSubscribableChannel[T interface{}](handler message.Handler[T]) *SubscribableChannel[T] {
	handlers := make([]message.Handler[T], 0)
	handlers = append(handlers, handler)
	return &SubscribableChannel[T]{MessageHandler: handlers}
}

// Send TODO HGA: CONSIDER TIMEOUT CONSTRAINTS
func (chann *SubscribableChannel[T]) Send(message message.Message[T], timeout int64) error {
	if message.GetPayload() == nil {
		return wserror.IllegalArgument{Message: "Message payload must not be null"}
	}
	return nil
}
