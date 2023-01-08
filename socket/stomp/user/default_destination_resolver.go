package user

import "mem-ws/socket/stomp/stompmsg"

type DefaultDestinationResolver[P interface{}, T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[P, T]) ResolveDestination(message stompmsg.Message[P]) DestinationResult {
	return DestinationResult{}
}
