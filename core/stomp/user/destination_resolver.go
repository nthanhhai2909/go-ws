package user

import (
	"mem-ws/core/message"
)

type DestinationResolver[P interface{}, T interface{}] interface {
	ResolveDestination(message message.Message[P, T]) DestinationResult
}
