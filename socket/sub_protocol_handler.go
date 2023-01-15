package socket

import (
	"mem-ws/socket/core/stomp/smsg"
)

// ISubProtocolHandler is specifications that all sub-protocol such as Stomp have to implement
type ISubProtocolHandler interface {

	// SupportProtocols - Get protocol supports list
	SupportProtocols() []string

	// HandleMessageFromClient channel is InboundChannel which process message from client
	HandleMessageFromClient(session IWebsocketSession, message WebsocketMessage)

	// SendMessageToClient - Send message to client
	SendMessageToClient(session IWebsocketSession, message smsg.IMessage[[]byte])
}
