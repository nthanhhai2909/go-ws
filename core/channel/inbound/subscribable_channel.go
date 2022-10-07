package inbound

import (
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/core/message"
	"mem-ws/core/simp"
	"mem-ws/core/simp/user"
	"mem-ws/core/wserror"
)

type subscribableChannel[T interface{}] struct {
	OutBoundChannels    map[string][]message.Handler[interface{}]
	ConnectChan         chan message.Handler[interface{}]
	DisConnectChan      chan message.Handler[interface{}]
	DestinationResolver user.DestinationResolver[T]
	//	Unregister  chan *core.Client
	//	Subscribe   chan *core.Subscribe
	//	Unsubscribe chan *core.Subscribe
	//	Topics      map[string]map[*core.Client]struct{}
	//	Broadcast   chan *core.Broadcast
}

// NewSubscribableChannel The only expose API to get SubscribableChannel instance
func NewSubscribableChannel() InboundChannel[interface{}] {
	chann := &subscribableChannel[interface{}]{
		OutBoundChannels: make(map[string][]message.Handler[interface{}], 0),
		ConnectChan:      make(chan message.Handler[interface{}]),
		DisConnectChan:   make(chan message.Handler[interface{}]),
	}
	go chann.startInternal()
	return chann
}

func (chann *subscribableChannel[T]) Connect(conn *websocket.Conn) (message.Handler[interface{}], error) {
	if conn == nil {
		return nil, wserror.IllegalArgument{Message: "Connection must not be null"}
	}
	handler := simp.NewBrokerMessageHandler(conn)
	chann.ConnectChan <- handler
	return handler, nil
}

func (chann *subscribableChannel[T]) Disconnect(handler message.Handler[interface{}]) {
	chann.DisConnectChan <- handler
}

func (chann *subscribableChannel[T]) startInternal() {
	for {
		select {
		case conn := <-chann.ConnectChan:
			chann.doConnectInternal(conn)
		case conn := <-chann.DisConnectChan:
			chann.doDisConnectInternal(conn)
			//case subscribe := <-h.Subscribe:
			//	h.subscribe(subscribe)
			//case subscribe := <-h.Unsubscribe:
			//	h.unsubscribe(subscribe)
			//case broadcast := <-h.Broadcast:
			//	h.broadcast(broadcast)
		}
	}
}

func (chann *subscribableChannel[T]) doConnectInternal(handler message.Handler[interface{}]) {
	fmt.Printf("New client %s connected", handler.GetUserID())
	chann.OutBoundChannels[handler.GetUserID()] = []message.Handler[interface{}]{handler}
}

func (chann *subscribableChannel[T]) doDisConnectInternal(handler message.Handler[interface{}]) {
	fmt.Printf("Client %s disconnected", handler.GetUserID())
	delete(chann.OutBoundChannels, handler.GetUserID())
	err := handler.GetConn().Close()
	close(handler.GetOutboundChannel())
	if err != nil {
		fmt.Printf("Error when close connect")
	}
}
