package native

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket/msg/types"
	"net/http"
)

type WebsocketSession struct {
	id               string
	acceptedProtocol string
	conn             *websocket.Conn
	outbound         chan []byte
	textSize         int
	binarySize       int
}

func NewWebsocketSession(conn *websocket.Conn, textSize int, binarySize int) *WebsocketSession {
	session := &WebsocketSession{
		conn:             conn,
		outbound:         make(chan []byte),
		id:               uuid.New().String(),
		acceptedProtocol: "TODO HGA",
		textSize:         textSize,
		binarySize:       binarySize,
	}

	go session.outboundInternal()
	return session
}

func (session *WebsocketSession) GetID() string {
	return session.id
}

func (session *WebsocketSession) GetHandshakeHeaders() http.Header {
	return nil
}

func (session *WebsocketSession) GetRemoteAddress() {
}

func (session *WebsocketSession) GetLocalAddress() {
}

func (session *WebsocketSession) GetAcceptedProtocol() string {
	return session.acceptedProtocol
}

func (session *WebsocketSession) SetTextMessageSizeLimit(size int) {
	session.textSize = size
}

func (session *WebsocketSession) GetTextMessageSizeLimit() int {
	return session.textSize
}

func (session *WebsocketSession) SetBinaryMessageSizeLimit(size int) {
	session.binarySize = size
}

func (session *WebsocketSession) GetBinaryMessageSizeLimit() int {
	return session.binarySize
}

func (session *WebsocketSession) GetExtensions() {
}

func (session *WebsocketSession) SendMessage(message types.WebsocketMessage) {
	session.outbound <- message.GetPayload()
}

func (session *WebsocketSession) IsOpen() bool {
	return false
}

func (session *WebsocketSession) Close() error {
	err := session.conn.Close()
	return err
}

// outboundInternal - process messagetype
func (session *WebsocketSession) outboundInternal() {

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
