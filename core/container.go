package core

import (
	"fmt"
	"log"
	"net/http"
)

type WSContainer interface {
	Handler(w http.ResponseWriter, r *http.Request)
	StartServe()
}

type wscontainer struct {
	websocketConnectionFactory *WebsocketConnectionFactory
	hub                        *Hub
}

func NewWSContainer(websocketConnectionFactory *WebsocketConnectionFactory) WSContainer {
	hub := NewHub()
	return &wscontainer{
		websocketConnectionFactory: websocketConnectionFactory,
		hub:                        hub,
	}
}

func (container *wscontainer) StartServe() {
	go container.hub.Start()
}

func (container *wscontainer) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := container.websocketConnectionFactory.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := NewClient(conn, container.hub)
	defer func() {
		err := client.Conn.Close()
		container.hub.Unregister <- client
		if err != nil {
			log.Println("Error when close connection")
		}
	}()

	client.Hub.Register <- client
	for {
		var req WSRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println("Error when read json: ", err)
			return
		}

		switch req.Action {
		case SUBSCRIBE:
			client.Subscribe(req)
		case UNSUBSCRIBE:
			client.Unsubscribe(req)
		case BROADCAST:
			client.Broadcast(req)
		case SEND_TO_USER:
			client.SendToUser(req)
		default:
			fmt.Println("Action do not support")
		}
	}
}
