package socket

import (
	"net/http"
)

// IWebsocketSession refer to a native Websocket Session established between client-server
type IWebsocketSession interface {
	GetID() string
	GetHandshakeHeaders() http.Header
	GetRemoteAddress()
	GetLocalAddress()
	GetAcceptedProtocol() string
	SetTextMessageSizeLimit(size int)
	GetTextMessageSizeLimit() int
	SetBinaryMessageSizeLimit(size int)
	GetBinaryMessageSizeLimit() int
	GetExtensions()
	SendMessage(message WebsocketMessage)
	IsOpen() bool
	Close() error
}
