package socket

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/socket/channel"
	"mem-ws/socket/socketmsg"
)

// SubProtocolWebsocketHandler - socket.WebSocketHandler & stompmsg.Handler Implementation
type SubProtocolWebsocketHandler struct {
	ClientInboundChannel channel.Channel
	SubProtocolHandler   ISubProtocolHandler
	Sessions             map[string]IWebsocketSession
}

func (h *SubProtocolWebsocketHandler) AfterConnectionEstablished(session IWebsocketSession) error {
	sessionId := session.GetID()
	if stringutils.IsBlank(sessionId) {
		return argument.Create("Session ID must not be null")
	}
	h.Sessions[sessionId] = session
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleMessageFromClient(session IWebsocketSession, message socketmsg.WebsocketMessage[[]byte]) error {
	h.SubProtocolHandler.HandleMessageFromClient(session, message, h.ClientInboundChannel)
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleTransportError(session IWebsocketSession, err error) error {
	return nil
}

func (h *SubProtocolWebsocketHandler) AfterConnectionClosed(session IWebsocketSession, status CloseStatus) error {
	return session.Close()
}

// TODO HGA WILL RESEARCH LATER
func (h *SubProtocolWebsocketHandler) SupportsPartialMessages() bool {
	return false
}
