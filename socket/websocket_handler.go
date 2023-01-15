package socket

// IWebsocketHandler take responsibility for handle message is sent from client
type IWebsocketHandler interface {
	AfterConnectionEstablished(session IWebsocketSession) error
	HandleMessageFromClient(session IWebsocketSession, message WebsocketMessage) error
	HandleTransportError(session IWebsocketSession, err error) error
	AfterConnectionClosed(session IWebsocketSession, status CloseStatus) error
	SupportsPartialMessages() bool
}
