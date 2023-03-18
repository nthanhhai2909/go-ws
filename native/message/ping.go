package message

type Ping struct {
	bytes []byte
}

func (msg *Ping) GetPayload() []byte {
	return msg.bytes
}

func (msg *Ping) GetPayloadLength() int {
	return len(msg.bytes)
}

// IsLast TODO HGA WILL UPDATE
func (msg *Ping) IsLast() bool {
	return true
}
