package socket

import (
	"mem-ws/socket/message"
	"net/http"
)

type WebsocketSession interface {
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
	SendMessage(message message.WebsocketMessage[[]byte])
	IsOpen() bool
	Close() error
}
