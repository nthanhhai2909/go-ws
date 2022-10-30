package stomp

import (
	"mem-ws/socket/channel"
	"mem-ws/socket/converter"
	"mem-ws/socket/stomp/msg"
	"mem-ws/socket/wserror"
)

type MessageTemplate[P interface{}] struct {
	MessageChannel   channel.Channel
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
	//return template.MessageChannel.Send(msg, template.Timeout)
	return nil
}
