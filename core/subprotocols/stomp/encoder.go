package stomp

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/core/subprotocols/stomp/constans"
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/message"
)

type Encoder struct {
}

func (e *Encoder) Encode(msg smsg.IMessage) message.IMessage {
	buff := bytes.NewBuffer(make([]byte, 0))
	headers := msg.GetMessageHeaders()
	command := headers.Command()
	buff.WriteString(fmt.Sprintf("%s\n", command))
	for key, value := range headers.Properties() {
		if key == header.Command {
			continue
		}
		buff.WriteString(fmt.Sprintf("%s:%s\n", key, value))
	}
	buff.WriteRune(constans.EndLineStringRune)
	if msg.GetPayload() != nil {
		buff.Write(msg.GetPayload())
		buff.WriteRune(constans.EndLineStringRune)
	}
	buff.WriteByte(constans.TerminalByte)
	return message.ToMessage(websocket.TextMessage, buff.Bytes())
}
