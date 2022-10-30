package socket

import "mem-ws/socket/stomp/msg"

type SendingOperations[P interface{}] interface {
	Send(destination string, message msg.Message[P]) error
}
