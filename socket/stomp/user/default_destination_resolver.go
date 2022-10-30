package user

import "mem-ws/socket/stomp/msg"

type DefaultDestinationResolver[P interface{}, T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[P, T]) ResolveDestination(message msg.Message[P]) DestinationResult {
	return DestinationResult{}
}
