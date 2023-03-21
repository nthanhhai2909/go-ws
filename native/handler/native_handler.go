package handler

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/native/enums"
	"mem-ws/native/message"
	"mem-ws/native/session"
	"mem-ws/native/subprotocol"
	"sync"
)

// NativeWebsocketHandler is used to support multiple Sub-protocol such as STOMP, AMQP, etc
type NativeWebsocketHandler struct {
	SubProtocolHandler subprotocol.ISubProtocolHandler
	Sessions           sync.Map
}

func (h *NativeWebsocketHandler) AfterConnectionEstablished(session session.ISession) error {
	sessionId := session.GetID()

	if stringutils.IsBlank(sessionId) {
		return argument.Create("Session ID must not be null")
	}
	h.Sessions.Store(sessionId, session)
	return nil
}

func (h *NativeWebsocketHandler) HandleMessageFromClient(session session.ISession, message message.IMessage) error {
	h.SubProtocolHandler.HandleMessageFromClient(session, message)
	return nil
}

func (h *NativeWebsocketHandler) HandleTransportError(session session.ISession, err error) error {
	return nil
}

func (h *NativeWebsocketHandler) AfterConnectionClosed(session session.ISession, status enums.CloseStatus) error {
	return session.Close()
}

// TODO HGA WILL RESEARCH LATER
func (h *NativeWebsocketHandler) SupportsPartialMessages() bool {
	return false
}
