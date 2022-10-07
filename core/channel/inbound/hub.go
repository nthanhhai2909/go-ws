package inbound

//type Hub struct {
//	Clients     map[*core.Client]struct{}
//	Register    chan *core.Client
//	Unregister  chan *core.Client
//	Subscribe   chan *core.Subscribe
//	Unsubscribe chan *core.Subscribe
//	Topics      map[string]map[*core.Client]struct{}
//	Broadcast   chan *core.Broadcast
//}
//
//func NewHub() *Hub {
//	return &Hub{
//		Clients:     make(map[*core.Client]struct{}),
//		Register:    make(chan *core.Client),
//		Unregister:  make(chan *core.Client),
//		Subscribe:   make(chan *core.Subscribe),
//		Unsubscribe: make(chan *core.Subscribe),
//		Topics:      make(map[string]map[*core.Client]struct{}),
//		Broadcast:   make(chan *core.Broadcast),
//	}
//}
//
//func (h *Hub) Start() {
//	for {
//		select {
//		case client := <-h.Register:
//			h.register(client)
//		case client := <-h.Unregister:
//			h.unregister(client)
//		case subscribe := <-h.Subscribe:
//			h.subscribe(subscribe)
//		case subscribe := <-h.Unsubscribe:
//			h.unsubscribe(subscribe)
//		case broadcast := <-h.Broadcast:
//			h.broadcast(broadcast)
//		}
//	}
//}
//
//func (h *Hub) register(client *core.Client) {
//	h.Clients[client] = struct{}{}
//}
//
//func (h *Hub) unregister(client *core.Client) {
//	fmt.Println("A user unregister")
//	if _, ok := h.Clients[client]; ok {
//		delete(h.Clients, client)
//		close(client.Outbound)
//		h.clearAllClientSub(client)
//	}
//}
//
//func (h *Hub) clearAllClientSub(client *core.Client) {
//	topics := client.SubscribeTopics
//	if len(topics) == 0 {
//		return
//	}
//
//	for _, topic := range topics {
//		h.unsubscribe(&core.Subscribe{Client: client, Destination: topic})
//	}
//}
//
//func (h *Hub) unsubscribe(subscribe *core.Subscribe) {
//	fmt.Println("A client unsubscribe to: ", subscribe.Destination)
//	client := subscribe.Client
//	destination := subscribe.Destination
//	client.DelSubscribeTopic(destination)
//	clients, ok := h.Topics[destination]
//	if !ok || len(clients) == 0 {
//		return
//	}
//
//	delete(clients, client)
//}
//
//func (h *Hub) broadcast(broadcast *core.Broadcast) {
//	sentClient := broadcast.Client
//	payload := broadcast.Data
//	destination := broadcast.Destination
//	if clients, ok := h.Topics[destination]; ok {
//		if _, existed := clients[sentClient]; existed {
//			for client, _ := range clients {
//				client.Outbound <- []byte(payload)
//			}
//		}
//	}
//}
//
//func (h *Hub) subscribe(subscribe *core.Subscribe) {
//	fmt.Println("A client subscribe to: ", subscribe.Destination)
//	client := subscribe.Client
//	destination := subscribe.Destination
//	client.AddSubscribeTopic(destination)
//	if clients, ok := h.Topics[destination]; ok {
//		if _, existed := clients[client]; !existed {
//			clients[client] = struct{}{}
//		}
//	} else {
//		clients := make(map[*core.Client]struct{})
//		clients[client] = struct{}{}
//		h.Topics[destination] = clients
//	}
//}
