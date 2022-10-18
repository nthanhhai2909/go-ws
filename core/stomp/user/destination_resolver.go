package user

import (
	"mem-ws/core/stomp/msg"
)

type DestinationResolver[P interface{}, T interface{}] interface {
	ResolveDestination(message msg.Message[P, T]) DestinationResult
}
