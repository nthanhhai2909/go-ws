package types

type PongMessage struct {
	bytes []byte
}

func (msg *PongMessage) GetPayload() []byte {
	return msg.bytes
}

func (msg *PongMessage) GetPayloadLength() int {
	return len(msg.bytes)
}

// IsLast TODO HGA WILL UPDATE
func (msg *PongMessage) IsLast() bool {
	return true
}
