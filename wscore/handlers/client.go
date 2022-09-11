package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"mem-ws/wscore/dto"
	"time"
)

const (
	writeWait = 10 * time.Second
)

type Subscribe struct {
	Client      *Client
	Destination string
}

type ClientBroadcast struct {
	Client      *Client
	Destination string
	Data        string
}

type Client struct {
	Conn           *websocket.Conn
	Hub            *Hub
	Outbound       chan []byte
	ID             string
	SubscribeTopic []string
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	client := &Client{
		Conn:           conn,
		Hub:            hub,
		Outbound:       make(chan []byte),
		ID:             uuid.New().String(),
		SubscribeTopic: make([]string, 0),
	}

	go client.outbound()
	return client
}

func (client *Client) AddSubscribeTopic(topic string) {
	client.SubscribeTopic = append(client.SubscribeTopic, topic)
}

func (client *Client) DelSubscribeTopic(topic string) {
	for index, ele := range client.SubscribeTopic {
		if ele == topic {
			client.SubscribeTopic = append(client.SubscribeTopic[:index], client.SubscribeTopic[index+1:]...)
		}
	}
}

func (client *Client) Broadcast(req dto.WSRequest) {
	var payload dto.UserBroadCast
	err := json.Unmarshal([]byte(req.Payload), &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	clientBroadcast := ClientBroadcast{Client: client, Destination: payload.Destination, Data: payload.Data}
	client.Hub.ClientBroadcast <- &clientBroadcast
}

func (client *Client) Unsubscribe(req dto.WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Unsubscribe <- &subscribe
}

func (client *Client) Subscribe(req dto.WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Subscribe <- &subscribe
}

func (client *Client) outbound() {
	fmt.Println("Client outbound started")
	defer func() {
		client.Conn.Close()
	}()
	for {
		select {
		case payload, ok := <-client.Outbound:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(payload)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
