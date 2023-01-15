package socket

import (
	"mem-ws/socket/core/stomp/smsg"
)

type SendingOperations[P interface{}] interface {
	Send(destination string, message smsg.IMessage) error
}
