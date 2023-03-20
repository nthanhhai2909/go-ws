package stomp

import (
	"bytes"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/cmd/client"
	"mem-ws/core/subprotocols/stomp/constans"
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
	"strings"
)

// Decoder TODO HGA WILL TEST LATER
type Decoder struct {
}

// Decode TODO HGA WILL PROCESS FOR ERROR MESSAGE
func (d *Decoder) Decode(buff []byte) (smsg.IMessage, error) {
	buffer := bytes.NewBuffer(buff)
	command, err := d.readCommand(buffer)
	if err != nil {
		return nil, err
	}
	headers, err := d.readHeaders(buffer)
	if err != nil {
		return nil, err
	}
	headers.AddHeader(header.Command, command.Type)
	payload, err := d.readPayload(buffer)
	if err != nil {
		return nil, err
	}
	return &smsg.GenericMessage[[]byte]{
		Headers: headers,
		Payload: payload,
	}, nil
}

func (d *Decoder) readCommand(buffer *bytes.Buffer) (*client.Command, error) {
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
	return client.ToCommand(comm.String()), nil
}

func (d *Decoder) readHeaders(buffer *bytes.Buffer) (*header.Header, error) {
	headers := header.EmptyHeader()
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
		headers.AddHeader(strs[0], strs[1])
		if isEnd, _ := d.tryToGetEndOfLine(buffer); isEnd {
			break
		}
	}
	return headers, nil
}

func (d *Decoder) readPayload(buffer *bytes.Buffer) ([]byte, error) {
	payload := make([]byte, 0)
	for {
		ch, err := buffer.ReadByte()
		if err != nil {
			return nil, err
		}
		if ch == constans.TerminalByte {
			break
		}
		payload = append(payload, ch)
	}
	return payload, nil
}

func (d *Decoder) tryToGetEndOfLine(buffer *bytes.Buffer) (bool, error) {
	if r, _, err := buffer.ReadRune(); err == nil {
		if r == '\n' {
			return true, nil
		} else if r == '\r' {
			if r, _, err := buffer.ReadRune(); err == nil && r == '\n' {
				return true, nil
			} else {
				return false, errors.MessageConversion{Message: "'\\r' must be followed by '\\n'"}
			}
		}
		err := buffer.UnreadRune()
		if err != nil {
			return false, err
		}
	}

	return false, nil
}
