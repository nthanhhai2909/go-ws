package wscore

import "fmt"

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() Hub {
	return Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			fmt.Println("New user connected")
			h.Clients[client] = true
		case client := <-h.Unregister:
			fmt.Println("User disconnected")
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Outbound)
			}
		case message := <-h.Broadcast:
			fmt.Println("Broadcast message to all user")
			for client := range h.Clients {
				select {
				case client.Outbound <- message:
				default:
					close(client.Outbound)
					delete(h.Clients, client)
				}
			}
		}
	}
}
