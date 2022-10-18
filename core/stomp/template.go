package stomp

import (
	"mem-ws/core/channel/outbound"
	"mem-ws/core/converter"
	"mem-ws/core/stomp/msg"
	"mem-ws/core/wserror"
)

type MessageTemplate[P interface{}] struct {
	MessageChannel   outbound.Channel[P]
	Timeout          int64
	MessageConverter converter.MessageConverter[P]
}

func (template *MessageTemplate[P]) Send(destination string, msg msg.Message[P]) error {
	if destination == "" {
		return wserror.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return wserror.IllegalArgument{Message: "Message is required"}
	}

	//msg.GetMessageHeaders().SetDestination(destination)
	return template.MessageChannel.Send(msg, template.Timeout)
}
