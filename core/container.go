package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/core/stomp"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wscontainer struct {
	websocketConnectionFactory *WebsocketConnectionFactory
	decoder                    *stomp.Decoder[interface{}, interface{}]
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	return &wscontainer{
		websocketConnectionFactory: websocketConnectionFactory,
		decoder:                    stomp.GetStompDecoder(),
	}
}

func (container *wscontainer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := container.websocketConnectionFactory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	inBoundChannel := container.websocketConnectionFactory.GetInboundChannel()
	handler, err := inBoundChannel.Connect(conn)
	if err != nil {
		log.Print("Create connection error:", err)
		return
	}

	defer func() {
		inBoundChannel.Disconnect(handler)
		if err != nil {
			log.Println("Error when close connection")
		}
	}()
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error when read message")
			return
		}
		if messageType != websocket.TextMessage {
			return
		}
		message := container.decoder.Decode(payload)
		fmt.Println(message)
	}
}
