package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/socket/stomp"
	"mem-ws/socket/stomp/cmd"
	"mem-ws/socket/stomp/msg"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wscontainer struct {
	factory *WebsocketConnectionFactory
	decoder *stomp.Decoder
	users   map[string]msg.Handler[[]byte]
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	return &wscontainer{
		factory: websocketConnectionFactory,
		decoder: stomp.GetStompDecoder(),
	}
}

func (container *wscontainer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := container.factory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	//inbound := container.factory.GetInboundChannel()
	//handler, err := inbound.Connect(conn)
	if err != nil {
		log.Println("Create connection error:", err)
		return
	}

	defer func() {
		//inbound.Disconnect(handler)
		if err != nil {
			log.Println("Error when close connection")
		}
	}()
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error when read msg")
			return
		}
		if messageType != websocket.TextMessage {
			return
		}
		message, _ := container.decoder.Decode(payload)
		headers := message.GetMessageHeaders()
		switch headers.GetCommand() {
		case cmd.Connect:
			//inbound.Connect()
		}
		if err != nil {
			log.Println("Error when send msg")
			return
		}
	}
}
