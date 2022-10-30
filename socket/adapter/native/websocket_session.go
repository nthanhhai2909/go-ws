package native

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"mem-ws/socket/message"
	"net/http"
	"time"
)

const (
	writeWait = 10 * time.Second
)

// WebsocketSession Implementation
type webSocketSession struct {
	id               string
	acceptedProtocol string
	conn             *websocket.Conn
	outbound         chan []byte
	textSize         int
	binarySize       int
}

func NewWebsocketSession(conn *websocket.Conn, textSize int, binarySize int) *webSocketSession {
	return &webSocketSession{
		conn:             conn,
		outbound:         make(chan []byte),
		id:               uuid.New().String(),
		acceptedProtocol: "TODO HGA",
		textSize:         textSize,
		binarySize:       binarySize,
	}
}

func (session *webSocketSession) GetID() string {
	return session.id
}

func (session *webSocketSession) GetHandshakeHeaders() http.Header {
	return nil
}

func (session *webSocketSession) GetRemoteAddress() {
}

func (session *webSocketSession) GetLocalAddress() {
}

func (session *webSocketSession) GetAcceptedProtocol() string {
	return session.acceptedProtocol
}

func (session *webSocketSession) SetTextMessageSizeLimit(size int) {
	session.textSize = size
}

func (session *webSocketSession) GetTextMessageSizeLimit() int {
	return session.textSize
}

func (session *webSocketSession) SetBinaryMessageSizeLimit(size int) {
	session.binarySize = size
}

func (session *webSocketSession) GetBinaryMessageSizeLimit() int {
	return session.binarySize
}

func (session *webSocketSession) GetExtensions() {
}

func (session *webSocketSession) SendMessage(message message.WebsocketMessage[[]byte]) {
	session.outbound <- message.GetPayload()
}

func (session *webSocketSession) IsOpen() bool {
	return false
}

func (session *webSocketSession) Close() error {
	err := session.conn.Close()
	return err
}

func (session *webSocketSession) listen() {

	defer func() {
		err := session.conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case payload, ok := <-session.outbound:
			err := session.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if !ok {
				err = session.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					// TODO HGA WILL ADAPT LATER
				}
				return
			}

			w, err := session.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(payload)
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
