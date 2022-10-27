package socket

import "net/http"

type WebsocketSession interface {
	GetID() string
	GetHandshakeHeaders() http.Header
	GetRemoteAddress()
	GetLocalAddress()
	GetAcceptedProtocol() string
	SetTextMessageSizeLimit(size int)
	GetTextMessageSizeLimit()
	SetBinaryMessageSizeLimit(size int)
	GetBinaryMessageSizeLimit()
	GetExtensions()
	SendMessage(message WebsocketMessage[interface{}])
	IsOpen() bool
	Close()
}
