package conf

import (
	"time"
)

type WsConfig struct {
	HandshakeTimeout time.Duration
	ReadBufferSize   int
	WriteBufferSize  int
	SubProtocols     []string
}
