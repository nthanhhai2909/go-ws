package core

import "mem-ws/core/message"

type SendingOperations interface {
	Send(destination string, message message.Message[interface{}]) error
}
