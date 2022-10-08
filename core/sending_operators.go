package core

import "mem-ws/core/message"

type SendingOperations[P interface{}, H interface{}] interface {
	Send(destination string, message message.Message[P, H]) error
}
