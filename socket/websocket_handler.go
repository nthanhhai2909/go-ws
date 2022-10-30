package socket

import (
	"mem-ws/socket/message"
)

type WebsocketHandler interface {
	AfterConnectionEstablished(session WebsocketSession) error
	HandleMessageFromClient(session WebsocketSession, message message.WebsocketMessage[[]byte]) error
	HandleTransportError(session WebsocketSession, err error) error
	AfterConnectionClosed(session WebsocketSession, status CloseStatus) error
	SupportsPartialMessages() bool
}
