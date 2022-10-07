package inbound

import (
	"github.com/gorilla/websocket"
	"mem-ws/core/message"
)

// InboundChannel TODO HGA WILL ADAPT
type InboundChannel[T interface{}] interface {
	Connect(conn *websocket.Conn) (message.Handler[interface{}], error)
	Disconnect(handler message.Handler[interface{}])
	//CloseConnection(message core.Handler[T]) error
	//Subscribe(destination string, message core.Handler[T]) error
	//Unsubscribe(destination string, message core.Handler[T]) error
}
