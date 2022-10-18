package msg

import (
	"github.com/gorilla/websocket"
)

type Handler[P interface{}] interface {
	HandleMessage(msg Message[P]) error
	GetUserID() string
	GetConn() *websocket.Conn
	GetOutboundChannel() chan []byte
}
