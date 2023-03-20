package smsg

import (
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/header"
)

type err struct {
	headers *header.Header
	payload []byte
}

func Error(headerProperties map[string]string, payload []byte) IMessage {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(header.Command, server.Error)
	for key, ele := range headerProperties {
		if header.Command == key {
			continue
		}
		headerInit.AddHeader(key, ele)
	}
	return &err{headers: headerInit, payload: payload}
}

func (msg *err) GetPayload() []byte {
	return msg.payload
}

func (msg *err) GetMessageHeaders() *header.Header {
	return msg.headers
}
