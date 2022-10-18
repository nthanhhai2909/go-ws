package inbound

import (
	"github.com/gorilla/websocket"
	message2 "mem-ws/core/stomp/msg"
)

// Channel TODO HGA WILL ADAPT
type Channel[H interface{}, P interface{}] interface {
	Connect(conn *websocket.Conn) (message2.Handler[P], error)
	Disconnect(handler message2.Handler[P])
	Subscribe(destination string, message message2.Handler[P]) error
	Unsubscribe(destination string, message message2.Handler[P]) error
	Send(message message2.Message[P]) error
}
