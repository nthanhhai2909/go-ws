package message

type Binary struct {
	bytes []byte
}

func (msg *Binary) GetPayload() []byte {
	return msg.bytes
}

func (msg *Binary) GetPayloadLength() int {
	return len(msg.bytes)
}

// IsLast TODO HGA WILL UPDATE
func (msg *Binary) IsLast() bool {
	return true
}
