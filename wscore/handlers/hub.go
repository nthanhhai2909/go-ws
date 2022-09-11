package handlers

type Hub struct {
	Clients     map[*Client]struct{}
	Register    chan *Client
	Unregister  chan *Client
	Subscribe   chan *Subscribe
	Unsubscribe chan *Subscribe
	Topics      map[string]map[*Client]struct{}
}

func NewHub() *Hub {
	return &Hub{
		Clients:     make(map[*Client]struct{}),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Subscribe:   make(chan *Subscribe),
		Unsubscribe: make(chan *Subscribe),
		Topics:      make(map[string]map[*Client]struct{}),
	}
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = struct{}{}
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Outbound)
			}
		case subscribe := <-h.Subscribe:
			h.subscribe(subscribe)
		case subscribe := <-h.Unsubscribe:
			h.unsubscribe(subscribe)
		}
	}
}

func (h *Hub) unsubscribe(subscribe *Subscribe) {
	clients, ok := h.Topics[subscribe.Destination]
	if !ok {
		return
	}

	delete(clients, subscribe.Client)
}

func (h *Hub) subscribe(subscribe *Subscribe) {
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
