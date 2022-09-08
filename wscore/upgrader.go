package wscore

import (
	"github.com/gorilla/websocket"
	"mem-ws/wscore/conf"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)

func NewInstance(conf conf.WsConfig) websocket.Upgrader {
	return websocket.Upgrader{
		HandshakeTimeout: conf.HandshakeTimeout,
		ReadBufferSize:   conf.ReadBufferSize,
		WriteBufferSize:  conf.WriteBufferSize,
		Subprotocols:     conf.SubProtocols,
	}
}

func DefInstance() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
	}
}
