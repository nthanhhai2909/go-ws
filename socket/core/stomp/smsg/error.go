package smsg

import (
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/cmd/server"
)

type err struct {
	headers *header.Headers
	payload []byte
}

func Error(headerProperties map[string]string, payload []byte) IMessage {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(header.CommandHeader, server.Error)
	for key, ele := range headerProperties {
		if header.CommandHeader == key {
			continue
		}
		headerInit.AddHeader(key, ele)
	}
	return &err{headers: headerInit, payload: payload}
}

func (msg *err) GetPayload() []byte {
	return msg.payload
}

func (msg *err) GetMessageHeaders() *header.Headers {
	return msg.headers
}
