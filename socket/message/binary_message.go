package message

type BinaryMessage struct {
	bytes []byte
}

func (msg *BinaryMessage) GetPayload() []byte {
	return msg.bytes
}

func (msg *BinaryMessage) GetPayloadLength() int {
	return len(msg.bytes)
}

// TODO HGA WILL UPDATE
func (msg *BinaryMessage) IsLast() bool {
	return true
}
