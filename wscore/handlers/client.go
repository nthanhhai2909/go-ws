package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
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
	Conn     *websocket.Conn
	Hub      *Hub
	Outbound chan []byte
	ID       string
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	client := &Client{
		Conn:     conn,
		Hub:      hub,
		Outbound: make(chan []byte),
		ID:       uuid.New().String(),
	}

	go client.outbound()
	return client
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
