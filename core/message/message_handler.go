package message

import (
	"github.com/gorilla/websocket"
)

type Handler[P interface{}, H interface{}] interface {
	HandleMessage(msg Message[P, H]) error
	GetUserID() string
	GetConn() *websocket.Conn
	GetOutboundChannel() chan []byte
}
