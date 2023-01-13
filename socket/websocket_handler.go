package socket

import "mem-ws/socket/msg/types"

// IWebsocketHandler take responsibility for handle message is sent from client
type IWebsocketHandler interface {
	AfterConnectionEstablished(session IWebsocketSession) error
	HandleMessageFromClient(session IWebsocketSession, message types.WebsocketMessage) error
	HandleTransportError(session IWebsocketSession, err error) error
	AfterConnectionClosed(session IWebsocketSession, status CloseStatus) error
	SupportsPartialMessages() bool
}
