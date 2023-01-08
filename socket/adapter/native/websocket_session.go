package native

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket/socketmsg"
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
	session := &webSocketSession{
		conn:             conn,
		outbound:         make(chan []byte),
		id:               uuid.New().String(),
		acceptedProtocol: "TODO HGA",
		textSize:         textSize,
		binarySize:       binarySize,
	}

	go session.startInternal()
	return session
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

func (session *webSocketSession) SendMessage(message socketmsg.WebsocketMessage[[]byte]) {
	session.outbound <- message.GetPayload()
}

func (session *webSocketSession) IsOpen() bool {
	return false
}

func (session *webSocketSession) Close() error {
	err := session.conn.Close()
	return err
}

func (session *webSocketSession) startInternal() {

	conn := session.conn
	defer func() {
		err := session.conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case payload, ok := <-session.outbound:
			//conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				log.Fatal("Error when send message to client")
			}
		}
	}
}
