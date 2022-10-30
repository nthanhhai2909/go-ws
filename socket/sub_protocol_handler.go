package socket

import (
	"mem-ws/socket/channel"
	"mem-ws/socket/stomp/msg"
)

type SubProtocolHandler interface {

	// SupportProtocols - Get protocol supports list
	SupportProtocols() []string

	// HandleMessageFromClient channel is InboundChannel which process message from client
	HandleMessageFromClient(session WebsocketSession, message msg.Message[[]byte], channel channel.Channel)

	// SendMessageToClient - Send message to client
	SendMessageToClient(session WebsocketSession, message msg.Message[[]byte])
}
