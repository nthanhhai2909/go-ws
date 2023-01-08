package socketmsg

type TextMessage struct {
	bytes []byte
}

func (msg *TextMessage) GetPayload() []byte {
	return msg.bytes
}

func (msg *TextMessage) GetPayloadLength() int {
	return len(msg.bytes)
}

// TODO HGA WILL UPDATE
func (msg *TextMessage) IsLast() bool {
	return true
}
