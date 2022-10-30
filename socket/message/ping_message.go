package message

type PingMessage struct {
	bytes []byte
}

func (msg *PingMessage) GetPayload() []byte {
	return msg.bytes
}

func (msg *PingMessage) GetPayloadLength() int {
	return len(msg.bytes)
}

// TODO HGA WILL UPDATE
func (msg *PingMessage) IsLast() bool {
	return true
}
