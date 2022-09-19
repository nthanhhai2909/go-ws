package simp

import error2 "mem-ws/core/error"

const (
	IndefiniteTimeout int = -1
)

type SimpleMessageTemplate[T interface{}] struct {
	MessageChannel Channel[T]
	Timeout        int64
}

func GetSimpleMessageTemplate[T interface{}](channel Channel[T], timeout int64) *SimpleMessageTemplate[T] {
	return &SimpleMessageTemplate[T]{
		MessageChannel: channel,
		Timeout:        timeout,
	}
}

func (template *SimpleMessageTemplate[T]) Send(destination string, msg Message[T]) error {
	if destination == "" {
		return error2.IllegalArgument{Message: "Destination is required"}
	}

	if msg == nil {
		return error2.IllegalArgument{Message: "Message is required"}
	}

	msg.GetMessageHeaders().SetDestination(destination)
	err := template.MessageChannel.Send(msg, template.Timeout)
	return err
}
