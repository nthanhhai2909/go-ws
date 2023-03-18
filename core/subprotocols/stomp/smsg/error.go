package smsg

import (
	"mem-ws/core/header"
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/constans"
)

type err struct {
	headers *header.Headers
	payload []byte
}

func Error(headerProperties map[string]string, payload []byte) IMessage {
	headerInit := header.EmptyHeader()
	headerInit.AddHeader(constans.CommandHeader, server.Error)
	for key, ele := range headerProperties {
		if constans.CommandHeader == key {
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
