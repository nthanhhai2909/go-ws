package outbound

import (
	"mem-ws/core/message"
)

type SubscribableChannel[P interface{}, H interface{}] struct {
	MessageHandler []message.Handler[P, H]
}

func NewSubscribableChannel(handler []message.Handler[interface{}, interface{}]) *SubscribableChannel[interface{}, interface{}] {
	return &SubscribableChannel[interface{}, interface{}]{MessageHandler: handler}
}

func NewSingleSubscribableChannel[P interface{}, H interface{}](handler message.Handler[P, H]) *SubscribableChannel[P, H] {
	handlers := make([]message.Handler[P, H], 0)
	handlers = append(handlers, handler)
	return &SubscribableChannel[P, H]{MessageHandler: handlers}
}

// Send TODO HGA: CONSIDER TIMEOUT CONSTRAINTS
func (chann *SubscribableChannel[P, T]) Send(message message.Message[P, T], timeout int64) error {
	return nil
}
