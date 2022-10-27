package msg

import (
	"github.com/gorilla/websocket"
)

type Handler[P interface{}] interface {
	HandleMessage(msg Message[P]) error
	GetConn() *websocket.Conn
	GetOutboundChannel() chan []byte
}
