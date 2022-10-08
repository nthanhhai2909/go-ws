package user

import (
	"mem-ws/core/message"
)

type DefaultDestinationResolver[P interface{}, T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[P, T]) ResolveDestination(message message.Message[P, T]) DestinationResult {
	return DestinationResult{}
}
