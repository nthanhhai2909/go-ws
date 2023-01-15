package user

import (
	"mem-ws/socket/core/stomp/smsg"
)

type DefaultDestinationResolver[P interface{}, T interface{}] struct {
}

// ResolveDestination TODO HGA WILL ADAPT LATER
func (resolver *DefaultDestinationResolver[P, T]) ResolveDestination(message smsg.IMessage[P]) DestinationResult {
	return DestinationResult{}
}
