package stompmsg

import (
	"mem-ws/socket/header"
	"mem-ws/socket/stomp/cmd/server"
)

type MessageBuilder struct{}

func (b *MessageBuilder) ConnectedMessage() Message[[]byte] {
	message := &GenericMessage[[]byte]{}
	messageHeaders := header.NewHeader()
	messageHeaders.SetHeader(header.CommandHeader, server.Connected)
	// TODO HGA CHECK CLIENT VERSION
	messageHeaders.SetHeader(header.StompVersionHeader, "1.1")
	messageHeaders.SetHeader(header.StompContentLengthHeader, "0")
	message.Headers = messageHeaders
	return message
}
