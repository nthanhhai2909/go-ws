package adapter

import (
	"github.com/gorilla/websocket"
	"mem-ws/socket"
	"net/http"
)

type StandardWebSocketSession struct {
	ID               string
	AcceptedProtocol string
	Conn             *websocket.Conn
}

func (session *StandardWebSocketSession) GetID() string {
	return session.ID
}

func (session *StandardWebSocketSession) GetHandshakeHeaders() http.Header {
	return nil
}

func (session *StandardWebSocketSession) GetRemoteAddress() {
}

func (session *StandardWebSocketSession) GetLocalAddress() {
}

func (session *StandardWebSocketSession) GetAcceptedProtocol() string {
	return "todo"
}

func (session *StandardWebSocketSession) SetTextMessageSizeLimit(size int) {
}

func (session *StandardWebSocketSession) GetTextMessageSizeLimit() int {
	return 0
}

func (session *StandardWebSocketSession) SetBinaryMessageSizeLimit(size int) {
}

func (session *StandardWebSocketSession) GetBinaryMessageSizeLimit() int {
	return 0
}

func (session *StandardWebSocketSession) GetExtensions() {
}

func (session *StandardWebSocketSession) SendMessage(message socket.WebsocketMessage[interface{}]) {
}

func (session *StandardWebSocketSession) IsOpen() bool {
	return false
}

func (session *StandardWebSocketSession) Close() {
}
