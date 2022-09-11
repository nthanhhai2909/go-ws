package handlers

import "fmt"

type Hub struct {
	Clients         map[*Client]struct{}
	Register        chan *Client
	Unregister      chan *Client
	Subscribe       chan *Subscribe
	Unsubscribe     chan *Subscribe
	Topics          map[string]map[*Client]struct{}
	ClientBroadcast chan *ClientBroadcast
}

func NewHub() *Hub {
	return &Hub{
		Clients:         make(map[*Client]struct{}),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		Subscribe:       make(chan *Subscribe),
		Unsubscribe:     make(chan *Subscribe),
		Topics:          make(map[string]map[*Client]struct{}),
		ClientBroadcast: make(chan *ClientBroadcast),
	}
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			h.register(client)
		case client := <-h.Unregister:
			h.unregister(client)
		case subscribe := <-h.Subscribe:
			h.subscribe(subscribe)
		case subscribe := <-h.Unsubscribe:
			h.unsubscribe(subscribe)
		case clientBroadcast := <-h.ClientBroadcast:
			h.clientBroadcast(clientBroadcast)
		}
	}
}

func (h *Hub) register(client *Client) {
	h.Clients[client] = struct{}{}
}

func (h *Hub) unregister(client *Client) {
	if _, ok := h.Clients[client]; ok {
		delete(h.Clients, client)
		close(client.Outbound)
	}
}

func (h *Hub) unsubscribe(subscribe *Subscribe) {
	fmt.Println("A client unsubscribe to: ", subscribe.Destination)
	clients, ok := h.Topics[subscribe.Destination]
	if !ok || len(clients) == 0 {
		return
	}

	delete(clients, subscribe.Client)
}

func (h *Hub) clientBroadcast(broadcast *ClientBroadcast) {
	sentClient := broadcast.Client
	payload := broadcast.Data
	destination := broadcast.Destination
	if clients, ok := h.Topics[destination]; ok {
		if _, existed := clients[sentClient]; existed {
			for client, _ := range clients {
				client.Outbound <- []byte(payload)
			}
		}
	}
}

func (h *Hub) subscribe(subscribe *Subscribe) {
	fmt.Println("A client subscribe to: ", subscribe.Destination)
	if clients, ok := h.Topics[subscribe.Destination]; ok {
		if _, existed := clients[subscribe.Client]; !existed {
			clients[subscribe.Client] = struct{}{}
		}
	} else {
		clients := make(map[*Client]struct{})
		clients[subscribe.Client] = struct{}{}
		h.Topics[subscribe.Destination] = clients
	}
}
