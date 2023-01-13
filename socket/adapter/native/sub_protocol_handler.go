package native

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/socket"
	"mem-ws/socket/core/channel"
	"mem-ws/socket/msg/types"
)

// SubProtocolWebsocketHandler is used to support multiple Sub-protocol such as STOMP, AMQP, etc
type SubProtocolWebsocketHandler struct {
	ClientInboundChannel channel.Channel
	SubProtocolHandler   socket.ISubProtocolHandler
	Sessions             map[string]socket.IWebsocketSession
}

func (h *SubProtocolWebsocketHandler) AfterConnectionEstablished(session socket.IWebsocketSession) error {
	sessionId := session.GetID()
	if stringutils.IsBlank(sessionId) {
		return argument.Create("Session ID must not be null")
	}
	h.Sessions[sessionId] = session
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleMessageFromClient(session socket.IWebsocketSession, message types.WebsocketMessage) error {
	h.SubProtocolHandler.HandleMessageFromClient(session, message, h.ClientInboundChannel)
	return nil
}

func (h *SubProtocolWebsocketHandler) HandleTransportError(session socket.IWebsocketSession, err error) error {
	return nil
}

func (h *SubProtocolWebsocketHandler) AfterConnectionClosed(session socket.IWebsocketSession, status socket.CloseStatus) error {
	return session.Close()
}

// TODO HGA WILL RESEARCH LATER
func (h *SubProtocolWebsocketHandler) SupportsPartialMessages() bool {
	return false
}
