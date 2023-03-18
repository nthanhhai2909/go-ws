package session

import (
	"mem-ws/native/message"
	"net/http"
)

type ISession interface {
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
	SendMessage(message message.IMessage)
	IsOpen() bool
	Close() error
}
