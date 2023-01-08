package socket

import (
	"mem-ws/socket/socketmsg"
	"net/http"
)

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
	SendMessage(message socketmsg.WebsocketMessage[[]byte])
	IsOpen() bool
	Close() error
}
