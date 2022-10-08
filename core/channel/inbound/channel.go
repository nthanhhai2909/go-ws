package inbound

import (
	"github.com/gorilla/websocket"
	"mem-ws/core/message"
)

// Channel TODO HGA WILL ADAPT
type Channel[P interface{}, H interface{}] interface {
	Connect(conn *websocket.Conn) (message.Handler[P, H], error)
	Disconnect(handler message.Handler[P, H])
	Subscribe(destination string, message message.Handler[P, H]) error
	Unsubscribe(destination string, message message.Handler[P, H]) error
	Send(message message.Message[P, H]) error
}
