package stomp

import (
	"mem-ws/socket/core/channel"
	"mem-ws/socket/core/converter"
	"mem-ws/socket/core/errors"
	"mem-ws/socket/core/stomp/stompmsg"
)

type MessageTemplate[P interface{}] struct {
	MessageChannel   channel.Channel
	Timeout          int64
	MessageConverter converter.MessageConverter[P]
}

func (template *MessageTemplate[P]) Send(destination string, msg stompmsg.Message[P]) error {
	if destination == "" {
		return errors.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return errors.IllegalArgument{Message: "Message is required"}
	}

	//stompmsg.GetMessageHeaders().SetDestination(destination)
	//return template.MessageChannel.Send(stompmsg, template.Timeout)
	return nil
}
