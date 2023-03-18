package message

type Pong struct {
	bytes []byte
}

func (msg *Pong) GetPayload() []byte {
	return msg.bytes
}

func (msg *Pong) GetPayloadLength() int {
	return len(msg.bytes)
}

// IsLast TODO HGA WILL UPDATE
func (msg *Pong) IsLast() bool {
	return true
}
