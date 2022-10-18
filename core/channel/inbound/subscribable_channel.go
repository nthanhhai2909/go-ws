package inbound

import (
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/core/stomp"
	message2 "mem-ws/core/stomp/msg"
	"mem-ws/core/wserror"
)

type subscribableChannel[T interface{}] struct {
	OutBoundChannels map[string][]message2.Handler[interface{}]
	ConnectChan      chan message2.Handler[interface{}]
	DisConnectChan   chan message2.Handler[interface{}]
	//	Unregister  chan *core.Client
	//	Subscribe   chan *core.Subscribe
	//	Unsubscribe chan *core.Subscribe
	//	Topics      map[string]map[*core.Client]struct{}
	//	Broadcast   chan *core.Broadcast
}

// NewSubscribableChannel The only expose API to get SubscribableChannel instance
func NewSubscribableChannel() Channel[interface{}, interface{}] {
	chann := &subscribableChannel[interface{}]{
		OutBoundChannels: make(map[string][]message2.Handler[interface{}], 0),
		ConnectChan:      make(chan message2.Handler[interface{}]),
		DisConnectChan:   make(chan message2.Handler[interface{}]),
	}
	go chann.startInternal()
	return chann
}

func (chann *subscribableChannel[T]) Connect(conn *websocket.Conn) (message2.Handler[interface{}], error) {
	if conn == nil {
		return nil, wserror.IllegalArgument{Message: "Connection must not be null"}
	}
	handler := stomp.NewBrokerMessageHandler(conn)
	chann.ConnectChan <- handler
	return handler, nil
}

func (chann *subscribableChannel[T]) Disconnect(handler message2.Handler[interface{}]) {
	chann.DisConnectChan <- handler
}

func (chann *subscribableChannel[T]) Subscribe(destination string, message message2.Handler[T]) error {
	return nil
}

func (chann *subscribableChannel[T]) Unsubscribe(destination string, message message2.Handler[T]) error {
	return nil
}

func (chann *subscribableChannel[T]) Send(message message2.Message[T]) error {
	return nil
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

func (chann *subscribableChannel[T]) doConnectInternal(handler message2.Handler[interface{}]) {
	fmt.Printf("New client %s connected", handler.GetUserID())
	chann.OutBoundChannels[handler.GetUserID()] = []message2.Handler[interface{}]{handler}
}

func (chann *subscribableChannel[T]) doDisConnectInternal(handler message2.Handler[interface{}]) {
	fmt.Printf("Client %s disconnected", handler.GetUserID())
	delete(chann.OutBoundChannels, handler.GetUserID())
	err := handler.GetConn().Close()
	close(handler.GetOutboundChannel())
	if err != nil {
		fmt.Printf("Error when close connect")
	}
}

func (chann *subscribableChannel[T]) doSubscribeInternal() {
}
