package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Subscribe struct {
	Client      *Client
	Destination string
}

type Client struct {
	Conn     *websocket.Conn
	Hub      *Hub
	Outbound chan []byte
	ID       string
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		Conn:     conn,
		Hub:      hub,
		Outbound: make(chan []byte),
		ID:       uuid.New().String(),
	}
}
