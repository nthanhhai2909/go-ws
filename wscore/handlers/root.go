package handlers

import (
	"encoding/json"
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
	defer handler.onConnectClosed(client)
	handler.onConnectionOpen(client)
	handler.onMessage(client)
}

func (handler *wshandler) onMessage(client *Client) {
	conn := client.Conn
	for {
		var req dto.WSRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		switch req.Action {
		case dto.SUBSCRIBE:
			subscribe(client, req)
		case dto.UNSUBSCRIBE:
			unsubscribe(client, req)
		case dto.BROADCAST:
			broadcast(client, req)
		default:
			fmt.Println("Action do not support")
		}
	}
}

func broadcast(client *Client, req dto.WSRequest) {
	var payload dto.UserBroadCast
	err := json.Unmarshal([]byte(req.Payload), &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	clientBroadcast := ClientBroadcast{Client: client, Destination: payload.Destination, Data: payload.Data}
	client.Hub.ClientBroadcast <- &clientBroadcast
}

func unsubscribe(client *Client, req dto.WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Unsubscribe <- &subscribe
}

func subscribe(client *Client, req dto.WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Subscribe <- &subscribe
}

func (handler *wshandler) onConnectionOpen(client *Client) {
	handler.hub.Register <- client
}

func (handler *wshandler) onConnectClosed(client *Client) {
	defer func() {
		err := client.Conn.Close()
		handler.hub.Unregister <- client
		if err != nil {
			log.Println("Error when close connection")
		}
	}()
}
