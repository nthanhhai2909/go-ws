package inbound

import (
	"mem-ws/socket/stomp/cmd"
	"mem-ws/socket/stomp/msg"
	"mem-ws/socket/wserror"
)

type SubscribableChannel struct {
	Handlers map[msg.Handler[[]byte]]struct{}
}

func NewSubscribableChannel() Channel[[]byte] {
	chann := &SubscribableChannel{
		Handlers: make(map[msg.Handler[[]byte]]struct{}, 0),
	}
	go chann.startInternal()
	return chann
}

//func (chann *SubscribableChannel) Subscribe(conn *websocket.Conn) (msg.Handler[[]byte], error) {
//	if conn == nil {
//		return nil, wserror.IllegalArgument{Message: "Connection must not be null"}
//	}
//	handler := stomp.NewBrokerMessageHandler(conn)
//	chann.ConnectChan <- handler
//	return handler, nil
//}
//
//func (chann *SubscribableChannel) Disconnect(handler msg.Handler[[]byte]) {
//	chann.DisConnectChan <- handler
//}

func (chann *SubscribableChannel) Subscribe(handler msg.Handler[[]byte]) error {
	return nil
}

func (chann *SubscribableChannel) Unsubscribe(handler msg.Handler[[]byte]) error {
	return nil
}

func (chann *SubscribableChannel) Send(message msg.Message[[]byte]) error {
	if message == nil {
		return wserror.IllegalArgument{Message: "Message must not be null"}
	}
	headers := message.GetMessageHeaders()
	switch headers.GetCommand() {
	case cmd.Connect:

	}
	return nil
}

func (chann *SubscribableChannel) startInternal() {
	for {
		select {
		//case conn := <-chann.ConnectChan:
		//	chann.doConnectInternal(conn)
		//case conn := <-chann.DisConnectChan:
		//	chann.doDisConnectInternal(conn)
		//case subscribe := <-h.Subscribe:
		//	h.subscribe(subscribe)
		//case subscribe := <-h.Unsubscribe:
		//	h.unsubscribe(subscribe)
		//case broadcast := <-h.Broadcast:
		//	h.broadcast(broadcast)
		}
	}
}

func (chann *SubscribableChannel) doConnectInternal(handler msg.Handler[[]byte]) {
	//fmt.Printf("New client %s connected", handler.GetUserID())
	//chann.OutBoundChannels[handler.GetUserID()] = []msg.Handler[[]byte]{handler}
}

func (chann *SubscribableChannel) doDisConnectInternal(handler msg.Handler[[]byte]) {
	//fmt.Printf("Client %s disconnected", handler.GetUserID())
	//delete(chann.OutBoundChannels, handler.GetUserID())
	//err := handler.GetConn().Close()
	//close(handler.GetOutboundChannel())
	//if err != nil {
	//	fmt.Printf("Error when close connect")
	//}
}

func (chann *SubscribableChannel) doSubscribeInternal() {
}
