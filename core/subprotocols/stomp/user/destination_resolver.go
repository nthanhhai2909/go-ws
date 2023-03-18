package user

import (
	"mem-ws/core/subprotocols/stomp/smsg"
)

type DestinationResolver[P interface{}] interface {
	ResolveDestination(message smsg.IMessage[P]) DestinationResult
}
