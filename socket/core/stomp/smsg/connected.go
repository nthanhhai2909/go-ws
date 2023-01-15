package smsg

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/cmd/server"
)

type connected struct {
	headers *header.Headers
}

func Connected(version string) IMessage[[]byte] {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(header.CommandHeader, server.Connected)
	headerInit.AddHeader(header.StompVersionHeader, version)
	return &connected{headers: headerInit}
}

func (msg *connected) GetPayload() []byte {
	return nil
}

func (msg *connected) GetMessageHeaders() *header.Headers {
	return msg.headers
}
