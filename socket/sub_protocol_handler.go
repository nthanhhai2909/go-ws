package socket

import "github.com/gorilla/websocket"

type SubProtocolHandler interface {
	SupportProtocols() []string

	HandlerMessageFromClient(conn *websocket.Conn)
}
