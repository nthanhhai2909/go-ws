package smsg

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/cmd/server"
)

type MessageBuilder struct{}

func (b *MessageBuilder) ConnectedMessage() Message[[]byte] {
	message := &GenericMessage[[]byte]{}
	messageHeaders := header.EmptyHeader()
	messageHeaders.AddHeader(header.CommandHeader, server.Connected)
	// TODO HGA CHECK CLIENT VERSION FROM CLIENT MESSAGE
	messageHeaders.AddHeader(header.StompVersionHeader, "1.1")
	messageHeaders.AddHeader(header.StompContentLengthHeader, "0")
	message.Headers = messageHeaders
	return message
}
