package user

import "mem-ws/socket/stomp/msg"

type DestinationResolver[P interface{}] interface {
	ResolveDestination(message msg.Message[P]) DestinationResult
}
