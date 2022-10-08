package stomp

import (
	"bytes"
	"fmt"
	"mem-ws/core/message"
)

// Decoder TODO HGA WILL TEST LATER
type Decoder[H interface{}, T interface{}] struct {
}

func GetStompDecoder() *Decoder[interface{}, interface{}] {
	return &Decoder[interface{}, interface{}]{}
}

func (d *Decoder[H, T]) Decode(payload []byte) message.Message[H, T] {
	buffer := bytes.NewBuffer(payload)
	command := d.readCommand(buffer)
	fmt.Println(command)
	return nil
}

func (d *Decoder[H, T]) readCommand(buffer *bytes.Buffer) Command {
	command := bytes.Buffer{}
	for !d.tryToGetEndOfLine(buffer) {
		ch, _ := buffer.ReadByte()
		command.WriteByte(ch)
	}
	fmt.Println(command.String())
	return ""
}

func (d *Decoder[H, T]) readHeaders() {

}

func (d *Decoder[H, T]) readPayload() {

}

func (d *Decoder[H, T]) tryToGetEndOfLine(buffer *bytes.Buffer) bool {
	ch, _ := buffer.ReadByte()
	if rune(ch) == '\n' {
		return true
	} else if rune(ch) == '\r' {
		if rune(ch) == '\n' {
			return true
		} else {
			return false
		}
	}
	return false
}
