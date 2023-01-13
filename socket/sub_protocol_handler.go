package socket

import (
	"mem-ws/socket/core/channel"
	"mem-ws/socket/core/stomp/stompmsg"
	"mem-ws/socket/msg/types"
)

type ISubProtocolHandler interface {

	// SupportProtocols - Get protocol supports list
	SupportProtocols() []string

	// HandleMessageFromClient channel is InboundChannel which process message from client
	HandleMessageFromClient(session IWebsocketSession, message types.WebsocketMessage, channel channel.Channel)

	// SendMessageToClient - Send message to client
	SendMessageToClient(session IWebsocketSession, message stompmsg.Message[[]byte])
}
