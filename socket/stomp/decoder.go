package stomp

import (
	"bytes"
	"mem-ws/socket/stomp/cmd"
	"mem-ws/socket/stomp/header"
	"mem-ws/socket/stomp/msg"
	"mem-ws/socket/wserror"
	"strings"
)

// Decoder TODO HGA WILL TEST LATER
type Decoder struct {
}

func GetStompDecoder() *Decoder {
	return &Decoder{}
}

// TODO HGA WILL PROCESS FOR ERROR MESSAGE
func (d *Decoder) Decode(buff []byte) (msg.Message[[]byte], error) {
	buffer := bytes.NewBuffer(buff)
	command, err := d.readCommand(buffer)
	if err != nil {
		return nil, err
	}
	headers, err := d.readHeaders(buffer)
	if err != nil {
		return nil, err
	}
	headers.SetCommand(command)
	payload, err := d.readPayload(buffer)
	if err != nil {
		return nil, err
	}
	return &msg.GenericMessage[[]byte]{
		Headers: headers,
		Payload: payload,
	}, nil
}

func (d *Decoder) readCommand(buffer *bytes.Buffer) (*cmd.Command, error) {
	comm := bytes.NewBuffer(make([]byte, 0))
	for {
		isEndLine, err := d.tryToGetEndOfLine(buffer)
		if isEndLine {
			break
		}

		if err != nil {
			return nil, err
		}

		ch, _, _ := buffer.ReadRune()
		comm.WriteRune(ch)
	}
	return cmd.ToCommand(comm.String()), nil
}

func (d *Decoder) readHeaders(buffer *bytes.Buffer) (*header.Headers, error) {
	headers := header.NewHeader()
	for {
		headerItem := bytes.NewBuffer(make([]byte, 0))
		for {
			isEndLine, err := d.tryToGetEndOfLine(buffer)
			if isEndLine {
				break
			}
			if err != nil {
				return nil, err
			}
			ch, _, _ := buffer.ReadRune()
			headerItem.WriteRune(ch)
		}
		strs := strings.Split(headerItem.String(), ":")
		headers.SetHeader(strs[0], strs[1])
		if isEnd, _ := d.tryToGetEndOfLine(buffer); isEnd {
			break
		}
	}
	return headers, nil
}

// TODO HGA WILL IMPLEMENT LATER
func (d *Decoder) readPayload(buffer *bytes.Buffer) ([]byte, error) {
	return nil, nil
}

func (d *Decoder) tryToGetEndOfLine(buffer *bytes.Buffer) (bool, error) {
	if r, _, err := buffer.ReadRune(); err == nil {
		if r == '\n' {
			return true, nil
		} else if r == '\r' {
			if r, _, err := buffer.ReadRune(); err == nil && r == '\n' {
				return true, nil
			} else {
				return false, wserror.MessageConversion{Message: "'\\r' must be followed by '\\n'"}
			}
		}
		err := buffer.UnreadRune()
		if err != nil {
			return false, err
		}
	}

	return false, nil
}
