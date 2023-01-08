package stomp

import (
	"mem-ws/socket"
	"mem-ws/socket/core/channel"
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/cmd/client"
	"mem-ws/socket/core/stomp/stompmsg"
)

// ProtocolHandler - socket.ISubProtocolHandler Implementation
type ProtocolHandler struct {
	Decoder        *Decoder
	Encoder        *Encoder
	MessageBuilder *stompmsg.MessageBuilder
}

func NewProtocolHandler() socket.ISubProtocolHandler {
	return &ProtocolHandler{
		Decoder:        GetStompDecoder(),
		Encoder:        GetStompEncoder(),
		MessageBuilder: &stompmsg.MessageBuilder{},
	}
}

func (h *ProtocolHandler) SupportProtocols() []string {
	return []string{"v10.stomp", "v11.stomp", "v12.stomp"}
}

func (h *ProtocolHandler) HandleMessageFromClient(session socket.IWebsocketSession, message socket.WebsocketMessage[[]byte], channel channel.Channel) {
	msg, _ := h.Decoder.Decode(message.GetPayload())
	headers := msg.GetMessageHeaders()
	switch headers.GetHeader(header.CommandHeader) {
	case client.Connect:
		h.SendMessageToClient(session, h.MessageBuilder.ConnectedMessage())
	}
}

func (h *ProtocolHandler) SendMessageToClient(session socket.IWebsocketSession, message stompmsg.Message[[]byte]) {
	session.SendMessage(h.Encoder.Encode(message))
}
