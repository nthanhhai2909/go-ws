package handler

import (
	error2 "mem-ws/native/enums"
	"mem-ws/native/message"
	"mem-ws/native/session"
)

// IWebsocketHandler take responsibility for handle message is sent from client
type IWebsocketHandler interface {
	AfterConnectionEstablished(session session.ISession) error
	HandleMessageFromClient(session session.ISession, message message.IMessage) error
	HandleTransportError(session session.ISession, err error) error
	AfterConnectionClosed(session session.ISession, status error2.CloseStatus) error
	SupportsPartialMessages() bool
}
