package wscore

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Outbound chan []byte
}
