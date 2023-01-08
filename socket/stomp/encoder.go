package stomp

import (
	"bytes"
	"github.com/gorilla/websocket"
	"mem-ws/socket/header"
	"mem-ws/socket/socketmsg"
	"mem-ws/socket/stomp/stompmsg"
)

type Encoder struct {
}

func GetStompEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Encode(msg stompmsg.Message[[]byte]) socketmsg.WebsocketMessage[[]byte] {
	msgBuffer := bytes.NewBuffer(make([]byte, 0))
	headers := msg.GetMessageHeaders()
	command := headers.GetHeader(header.CommandHeader)
	msgBuffer.WriteString(command)
	msgBuffer.WriteString("\n")
	for key, value := range headers.GetHeaderProperties() {
		if key == header.CommandHeader {
			continue
		}
		msgBuffer.WriteString(key)
		msgBuffer.WriteString(":")
		msgBuffer.WriteString(value)
		msgBuffer.WriteString("\n")
	}
	msgBuffer.WriteString("\n")
	if msg.GetPayload() != nil {
		msgBuffer.Write(msg.GetPayload())
		msgBuffer.WriteString("\n")
	}
	msgBuffer.Write([]byte{0})
	return socketmsg.ToWebsocketMessage(websocket.TextMessage, msgBuffer.Bytes())
}
