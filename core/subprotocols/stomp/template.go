package stomp

import (
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/channel"
	"mem-ws/core/subprotocols/stomp/converter"
	"mem-ws/core/subprotocols/stomp/smsg"
)

type MessageTemplate[P interface{}] struct {
	MessageChannel   channel.Channel
	Timeout          int64
	MessageConverter converter.MessageConverter[P]
}

func (template *MessageTemplate[P]) Send(destination string, msg smsg.IMessage) error {
	if destination == "" {
		return errors.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return errors.IllegalArgument{Message: "IMessage is required"}
	}

	//smsg.GetMessageHeaders().SetDestination(destination)
	//return template.MessageChannel.Send(smsg, template.Timeout)
	return nil
}
