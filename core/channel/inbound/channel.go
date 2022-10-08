package inbound

import (
	"github.com/gorilla/websocket"
	"mem-ws/core/message"
)

// Channel TODO HGA WILL ADAPT
type Channel[T interface{}] interface {
	Connect(conn *websocket.Conn) (message.Handler[T], error)
	Disconnect(handler message.Handler[T])
	Subscribe(destination string, message message.Handler[T]) error
	//Unsubscribe(destination string, message core.Handler[T]) error
}
