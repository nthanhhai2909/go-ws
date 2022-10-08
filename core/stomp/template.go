package stomp

import (
	"mem-ws/core/channel/outbound"
	"mem-ws/core/converter"
	"mem-ws/core/message"
	"mem-ws/core/wserror"
)

type MessageTemplate[P interface{}, H interface{}] struct {
	MessageChannel   outbound.Channel[P, H]
	Timeout          int64
	MessageConverter converter.MessageConverter[P, H]
}

func (template *MessageTemplate[P, T]) Send(destination string, msg message.Message[P, T]) error {
	if destination == "" {
		return wserror.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return wserror.IllegalArgument{Message: "Message is required"}
	}

	//msg.GetMessageHeaders().SetDestination(destination)
	return template.MessageChannel.Send(msg, template.Timeout)
}
