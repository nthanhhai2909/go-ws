package inbound

import (
	"fmt"
	"github.com/gorilla/websocket"
	"mem-ws/core/stomp"
	"mem-ws/core/stomp/cmd"
	"mem-ws/core/stomp/msg"
	"mem-ws/core/wserror"
)

type subscribableChannel struct {
	OutBoundChannels map[string][]msg.Handler[[]byte]
	ConnectChan      chan msg.Handler[[]byte]
	DisConnectChan   chan msg.Handler[[]byte]
	//	Unregister  chan *core.Client
	//	Subscribe   chan *core.Subscribe
	//	Unsubscribe chan *core.Subscribe
	//	Topics      map[string]map[*core.Client]struct{}
	//	Broadcast   chan *core.Broadcast
}

// NewSubscribableChannel The only expose API to get SubscribableChannel instance
func NewSubscribableChannel() Channel[[]byte] {
	chann := &subscribableChannel{
		OutBoundChannels: make(map[string][]msg.Handler[[]byte], 0),
		ConnectChan:      make(chan msg.Handler[[]byte]),
		DisConnectChan:   make(chan msg.Handler[[]byte]),
	}
	go chann.startInternal()
	return chann
}

func (chann *subscribableChannel) Connect(conn *websocket.Conn) (msg.Handler[[]byte], error) {
	if conn == nil {
		return nil, wserror.IllegalArgument{Message: "Connection must not be null"}
	}
	handler := stomp.NewBrokerMessageHandler(conn)
	chann.ConnectChan <- handler
	return handler, nil
}

func (chann *subscribableChannel) Disconnect(handler msg.Handler[[]byte]) {
	chann.DisConnectChan <- handler
}

func (chann *subscribableChannel) Subscribe(destination string, message msg.Handler[[]byte]) error {
	return nil
}

func (chann *subscribableChannel) Unsubscribe(destination string, message msg.Handler[[]byte]) error {
	return nil
}

func (chann *subscribableChannel) Send(message msg.Message[[]byte]) error {
	if message == nil {
		return wserror.IllegalArgument{Message: "Message must not be null"}
	}
	headers := message.GetMessageHeaders()
	switch headers.GetCommand() {
	case cmd.Connect:

	}
	return nil
}

func (chann *subscribableChannel) startInternal() {
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

func (chann *subscribableChannel) doConnectInternal(handler msg.Handler[[]byte]) {
	fmt.Printf("New client %s connected", handler.GetUserID())
	chann.OutBoundChannels[handler.GetUserID()] = []msg.Handler[[]byte]{handler}
}

func (chann *subscribableChannel) doDisConnectInternal(handler msg.Handler[[]byte]) {
	fmt.Printf("Client %s disconnected", handler.GetUserID())
	delete(chann.OutBoundChannels, handler.GetUserID())
	err := handler.GetConn().Close()
	close(handler.GetOutboundChannel())
	if err != nil {
		fmt.Printf("Error when close connect")
	}
}

func (chann *subscribableChannel) doSubscribeInternal() {
}
