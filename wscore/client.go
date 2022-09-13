package wscore

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait = 10 * time.Second
)

type Subscribe struct {
	Client      *Client
	Destination string
}

type Broadcast struct {
	Client      *Client
	Destination string `json:"destination"`
	Data        string `json:"data"`
}

type Client struct {
	Conn            *websocket.Conn
	Hub             *Hub
	Outbound        chan []byte
	ID              string
	SubscribeTopics []string
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	client := &Client{
		Conn:            conn,
		Hub:             hub,
		Outbound:        make(chan []byte),
		ID:              uuid.New().String(),
		SubscribeTopics: make([]string, 0),
	}

	go client.outbound()
	return client
}

func (client *Client) AddSubscribeTopic(topic string) {
	client.SubscribeTopics = append(client.SubscribeTopics, topic)
}

func (client *Client) DelSubscribeTopic(topic string) {
	for index, ele := range client.SubscribeTopics {
		if ele == topic {
			client.SubscribeTopics = append(client.SubscribeTopics[:index], client.SubscribeTopics[index+1:]...)
		}
	}
}

func (client *Client) Broadcast(req WSRequest) {
	var payload Broadcast
	err := json.Unmarshal([]byte(req.Payload), &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	payload.Client = client
	client.Hub.Broadcast <- &payload
}

func (client *Client) Unsubscribe(req WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Unsubscribe <- &subscribe
}

func (client *Client) Subscribe(req WSRequest) {
	payload := req.Payload
	subscribe := Subscribe{Client: client, Destination: payload}
	client.Hub.Subscribe <- &subscribe
}

func (client *Client) SendToUser(req WSRequest) {

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
