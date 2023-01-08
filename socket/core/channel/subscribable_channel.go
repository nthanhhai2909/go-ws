package channel

import (
	"mem-ws/socket/core/errors"
	stompmsg2 "mem-ws/socket/core/stomp/stompmsg"
)

type SubscribableChannel struct {
	Handlers map[stompmsg2.Handler]struct{}
}

func NewSubscribableChannel() Channel {
	chann := &SubscribableChannel{
		Handlers: make(map[stompmsg2.Handler]struct{}, 0),
	}
	go chann.startInternal()
	return chann
}

//func (chann *SubscribableChannel) Subscribe(conn *websocket.Conn) (stompmsg.Handler[[]byte], error) {
//	if conn == nil {
//		return nil, errors.IllegalArgument{Message: "Connection must not be null"}
//	}
//	handler := stomp.NewBrokerMessageHandler(conn)
//	chann.ConnectChan <- handler
//	return handler, nil
//}
//
//func (chann *SubscribableChannel) Disconnect(handler stompmsg.Handler[[]byte]) {
//	chann.DisConnectChan <- handler
//}

func (chann *SubscribableChannel) Subscribe(handler stompmsg2.Handler) error {
	return nil
}

func (chann *SubscribableChannel) Unsubscribe(handler stompmsg2.Handler) error {
	return nil
}

func (chann *SubscribableChannel) Send(message stompmsg2.Message[interface{}]) error {
	if message == nil {
		return errors.IllegalArgument{Message: "Message must not be null"}
	}
	//headers := message.GetMessageHeaders()
	//switch headers.GetCommand() {
	//case client.Connect:
	//
	//}
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

func (chann *SubscribableChannel) doConnectInternal(handler stompmsg2.Handler) {
	//fmt.Printf("New client %s connected", handler.GetUserID())
	//chann.OutBoundChannels[handler.GetUserID()] = []stompmsg.Handler[[]byte]{handler}
}

func (chann *SubscribableChannel) doDisConnectInternal(handler stompmsg2.Handler) {
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
