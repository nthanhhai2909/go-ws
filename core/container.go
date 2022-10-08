package core

import (
	"fmt"
	"log"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wscontainer struct {
	websocketConnectionFactory *WebsocketConnectionFactory
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	return &wscontainer{
		websocketConnectionFactory: websocketConnectionFactory,
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
		fmt.Println(messageType)
		fmt.Println(payload)
		if err != nil {
			log.Println("Error when read json: ", err)
			return
		}

		//switch req.Action {
		//case SUBSCRIBE:
		//	client.Subscribe(req)
		//case UNSUBSCRIBE:
		//	client.Unsubscribe(req)
		//case BROADCAST:
		//	client.Broadcast(req)
		//case SEND_TO_USER:
		//	client.SendToUser(req)
		//default:
		//	fmt.Println("Action do not support")
		//}
	}
}
