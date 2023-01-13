package stomp

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/socket/core/header"
	"mem-ws/socket/core/stomp/smsg"
	"mem-ws/socket/msg/types"
)

type Encoder struct {
}

func (e *Encoder) Encode(msg smsg.Message[[]byte]) types.WebsocketMessage {
	buff := bytes.NewBuffer(make([]byte, 0))
	headers := msg.GetMessageHeaders()
	command := headers.GetHeader(header.CommandHeader)
	buff.WriteString(fmt.Sprintf("%s\n", command))
	for key, value := range headers.Properties {
		if key == header.CommandHeader {
			continue
		}
		buff.WriteString(fmt.Sprintf("%s:%s\n", key, value))
	}
	buff.WriteRune(EndLineStringRune)
	if msg.GetPayload() != nil {
		buff.Write(msg.GetPayload())
		buff.WriteRune(EndLineStringRune)
	}
	buff.WriteByte(TerminalByte)
	return types.ToWebsocketMessage(websocket.TextMessage, buff.Bytes())
}
