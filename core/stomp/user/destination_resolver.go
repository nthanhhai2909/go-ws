package user

import (
	"mem-ws/core/message"
)

type DestinationResolver[T interface{}] interface {
	ResolveDestination(message message.Message[T]) DestinationResult
}
