package user

import (
	"mem-ws/socket/core/stomp/smsg"
)

type DestinationResolver[P interface{}] interface {
	ResolveDestination(message smsg.IMessage[P]) DestinationResult
}