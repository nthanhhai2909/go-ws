package socket

import (
	"mem-ws/socket/channel"
	"mem-ws/socket/socketmsg"
	"mem-ws/socket/stomp/stompmsg"
)

type ISubProtocolHandler interface {

	// SupportProtocols - Get protocol supports list
	SupportProtocols() []string

	// HandleMessageFromClient channel is InboundChannel which process socketmsg from client
	HandleMessageFromClient(session IWebsocketSession, message socketmsg.WebsocketMessage[[]byte], channel channel.Channel)

	// SendMessageToClient - Send socketmsg to client
	SendMessageToClient(session IWebsocketSession, message stompmsg.Message[[]byte])
}
