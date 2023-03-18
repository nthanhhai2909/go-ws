package message

type Text struct {
	bytes []byte
}

func (msg *Text) GetPayload() []byte {
	return msg.bytes
}

func (msg *Text) GetPayloadLength() int {
	return len(msg.bytes)
}

// IsLast TODO HGA WILL UPDATE
func (msg *Text) IsLast() bool {
	return true
}
