package user

import (
	"mem-ws/core/stomp/msg"
)

type DefaultDestinationResolver[P interface{}, T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[P, T]) ResolveDestination(message msg.Message[P, T]) DestinationResult {
	return DestinationResult{}
}
