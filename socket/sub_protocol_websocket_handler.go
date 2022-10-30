package socket

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/socket/channel"
	"mem-ws/socket/message"
)

// SubProtocolWebsocketHandler - socket.WebSocketHandler & msg.Handler Implementation
type SubProtocolWebsocketHandler struct {
	ClientInboundChannel channel.Channel
	SubProtocolHandler   SubProtocolHandler
	Sessions             map[string]WebsocketSession
}

func (h *SubProtocolWebsocketHandler) AfterConnectionEstablished(session WebsocketSession) error {
	sessionId := session.GetID()
	if stringutils.IsBlank(sessionId) {
		return argument.Create("Session ID must not be null")

	}
	h.Sessions[sessionId] = session
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleMessageFromClient(session WebsocketSession, message message.WebsocketMessage[[]byte]) error {
	h.SubProtocolHandler.HandleMessageFromClient(session, message, h.ClientInboundChannel)
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleTransportError(session WebsocketSession, err error) error {
	return nil
}

func (h *SubProtocolWebsocketHandler) AfterConnectionClosed(session WebsocketSession, status CloseStatus) error {
	return session.Close()
}

// TODO HGA WILL RESEARCH LATER
func (h *SubProtocolWebsocketHandler) SupportsPartialMessages() bool {
	return false
}
