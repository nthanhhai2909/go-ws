package user

import (
	"mem-ws/core/message"
)

type DefaultDestinationResolver[T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[T]) ResolveDestination(message message.Message[T]) DestinationResult {
	return DestinationResult{}
}
