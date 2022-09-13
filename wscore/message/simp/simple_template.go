package simp

import "mem-ws/wscore/message"

type SimpleMessageTemplate struct {
	messageChannel message.Channel
}

func (template *SimpleMessageTemplate) Send(destination string, msg message.Message[interface{}]) error {
	if destination == "" {
		return message.Error{Message: "Destination is required"}
	}

	if msg == nil {
		return message.Error{Message: "Message is required"}
	}

	return nil
}
