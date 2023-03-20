package smsg

import (
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/header"
)

type connected struct {
	headers *header.Header
}

func Connected(version string) IMessage {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(header.Command, server.Connected)
	headerInit.AddHeader(header.StompVersion, version)
	return &connected{headers: headerInit}
}

func (msg *connected) GetPayload() []byte {
	return nil
}

func (msg *connected) GetMessageHeaders() *header.Header {
	return msg.headers
}
