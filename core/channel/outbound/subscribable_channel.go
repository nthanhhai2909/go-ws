package outbound

import (
	message2 "mem-ws/core/stomp/msg"
)

type SubscribableChannel[P interface{}] struct {
	MessageHandler []message2.Handler[P]
}

func NewSubscribableChannel(handler []message2.Handler[interface{}]) *SubscribableChannel[interface{}] {
	return &SubscribableChannel[interface{}]{MessageHandler: handler}
}

func NewSingleSubscribableChannel[P interface{}](handler message2.Handler[P]) *SubscribableChannel[P] {
	handlers := make([]message2.Handler[P], 0)
	handlers = append(handlers, handler)
	return &SubscribableChannel[P]{MessageHandler: handlers}
}

// Send TODO HGA: CONSIDER TIMEOUT CONSTRAINTS
func (chann *SubscribableChannel[P]) Send(message message2.Message[P], timeout int64) error {
	return nil
}
