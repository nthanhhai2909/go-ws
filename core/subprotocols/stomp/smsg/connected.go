package smsg

import (
	"mem-ws/core/header"
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/constans"
)

type connected struct {
	headers *header.Headers
}

func Connected(version string) IMessage {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(constans.CommandHeader, server.Connected)
	headerInit.AddHeader(constans.StompVersionHeader, version)
	return &connected{headers: headerInit}
}

func (msg *connected) GetPayload() []byte {
	return nil
}

func (msg *connected) GetMessageHeaders() *header.Headers {
	return msg.headers
}
