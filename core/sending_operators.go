package core

import (
	"mem-ws/core/stomp/msg"
)

type SendingOperations[P interface{}] interface {
	Send(destination string, message msg.Message[P]) error
}
