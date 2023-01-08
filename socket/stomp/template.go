package stomp

import (
	"mem-ws/socket/channel"
	"mem-ws/socket/converter"
	"mem-ws/socket/stomp/stompmsg"
	"mem-ws/socket/wserror"
)

type MessageTemplate[P interface{}] struct {
	MessageChannel   channel.Channel
	Timeout          int64
	MessageConverter converter.MessageConverter[P]
}

func (template *MessageTemplate[P]) Send(destination string, msg stompmsg.Message[P]) error {
	if destination == "" {
		return wserror.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return wserror.IllegalArgument{Message: "Message is required"}
	}

	//socketmsg.GetMessageHeaders().SetDestination(destination)
	//return template.MessageChannel.Send(socketmsg, template.Timeout)
	return nil
}
