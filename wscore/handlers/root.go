package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"mem-ws/wscore/dto"
	"net/http"
)

type WSHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type wshandler struct {
	upg websocket.Upgrader
	hub *Hub
}

func New(upg websocket.Upgrader, hub *Hub) WSHandler {
	return &wshandler{
		upg: upg,
		hub: hub,
	}
}

func (handler *wshandler) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.upg.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := NewClient(conn, handler.hub)
	defer func() {
		err := client.Conn.Close()
		handler.hub.Unregister <- client
		if err != nil {
			log.Println("Error when close connection")
		}
	}()

	client.Hub.Register <- client
	for {
		var req dto.WSRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		switch req.Action {
		case dto.SUBSCRIBE:
			client.Subscribe(req)
		case dto.UNSUBSCRIBE:
			client.Unsubscribe(req)
		case dto.BROADCAST:
			client.Broadcast(req)
		default:
			fmt.Println("Action do not support")
		}
	}
}
