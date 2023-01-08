package user

import "mem-ws/socket/stomp/stompmsg"

type DestinationResolver[P interface{}] interface {
	ResolveDestination(message stompmsg.Message[P]) DestinationResult
}
