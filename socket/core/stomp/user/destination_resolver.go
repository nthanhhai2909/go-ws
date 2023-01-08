package user

import (
	"mem-ws/socket/core/stomp/stompmsg"
)

type DestinationResolver[P interface{}] interface {
	ResolveDestination(message stompmsg.Message[P]) DestinationResult
}
