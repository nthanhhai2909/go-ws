package message

import "github.com/gorilla/websocket"

func ToMessage(messageType int, payload []byte) IMessage {
	switch messageType {
	case websocket.TextMessage:
		return &Text{bytes: payload}
	case websocket.BinaryMessage:
		return &Binary{bytes: payload}
	case websocket.PingMessage:
		return &Ping{bytes: payload}
	case websocket.PongMessage:
		return &Pong{bytes: payload}
	}
	return nil
}
