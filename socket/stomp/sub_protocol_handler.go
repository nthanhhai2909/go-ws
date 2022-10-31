package stomp

import (
	"fmt"
	"mem-ws/socket"
	"mem-ws/socket/channel"
	"mem-ws/socket/message"
	"mem-ws/socket/stomp/msg"
)

// SubProtocolHandler - socket.SubProtocolHandler Implementation
type SubProtocolHandler struct {
}

func (h *SubProtocolHandler) SupportProtocols() []string {
	return []string{"v10.stomp", "v11.stomp", "v12.stomp"}
}

func (h *SubProtocolHandler) HandleMessageFromClient(session socket.WebsocketSession, message message.WebsocketMessage[[]byte], channel channel.Channel) {
	fmt.Println("HIHI NEK")
}

func (h *SubProtocolHandler) SendMessageToClient(session socket.WebsocketSession, message msg.Message[[]byte]) {
}
