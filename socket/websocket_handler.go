package socket

import (
	"mem-ws/socket/socketmsg"
)

type IWebsocketHandler interface {
	AfterConnectionEstablished(session IWebsocketSession) error
	HandleMessageFromClient(session IWebsocketSession, message socketmsg.WebsocketMessage[[]byte]) error
	HandleTransportError(session IWebsocketSession, err error) error
	AfterConnectionClosed(session IWebsocketSession, status CloseStatus) error
	SupportsPartialMessages() bool
}
