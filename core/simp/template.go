package simp

import (
	"mem-ws/core"
	"mem-ws/core/channel"
	"mem-ws/core/converter"
	"mem-ws/core/wserror"
)

type SimpleMessageTemplate[T interface{}] struct {
	MessageChannel   channel.Channel[T]
	Timeout          int64
	MessageConverter converter.MessageConverter[T]
}

func (template *SimpleMessageTemplate[T]) Send(destination string, msg core.Message[T]) error {
	if destination == "" {
		return wserror.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return wserror.IllegalArgument{Message: "Message is required"}
	}

	msg.GetMessageHeaders().SetDestination(destination)
	return template.MessageChannel.Send(msg, template.Timeout)
}
