package socket

import (
	"github.com/gorilla/websocket"
)

type WebsocketMessage interface {
	GetPayload() []byte
	GetPayloadLength() int
	IsLast() bool
}

func ToWebsocketMessage(messageType int, payload []byte) WebsocketMessage {
	switch messageType {
	case websocket.TextMessage:
		return &TextMessage{bytes: payload}
	case websocket.BinaryMessage:
		return &BinaryMessage{bytes: payload}
	case websocket.PingMessage:
		return &PingMessage{bytes: payload}
	case websocket.PongMessage:
		return &PongMessage{bytes: payload}
	}
	return nil
}
