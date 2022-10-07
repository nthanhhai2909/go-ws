package message

import (
	"github.com/gorilla/websocket"
)

type Handler[T interface{}] interface {
	HandleMessage(msg Message[T]) error
	GetUserID() string
	GetConn() *websocket.Conn
	GetOutboundChannel() chan []byte
}
