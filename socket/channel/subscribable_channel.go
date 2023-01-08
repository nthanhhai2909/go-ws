package channel

import (
	"mem-ws/socket/stomp/stompmsg"
	"mem-ws/socket/wserror"
)

type SubscribableChannel struct {
	Handlers map[stompmsg.Handler]struct{}
}

func NewSubscribableChannel() Channel {
	chann := &SubscribableChannel{
		Handlers: make(map[stompmsg.Handler]struct{}, 0),
	}
	go chann.startInternal()
	return chann
}

//func (chann *SubscribableChannel) Subscribe(conn *websocket.Conn) (socketmsg.Handler[[]byte], error) {
//	if conn == nil {
//		return nil, wserror.IllegalArgument{Message: "Connection must not be null"}
//	}
//	handler := stomp.NewBrokerMessageHandler(conn)
//	chann.ConnectChan <- handler
//	return handler, nil
//}
//
//func (chann *SubscribableChannel) Disconnect(handler socketmsg.Handler[[]byte]) {
//	chann.DisConnectChan <- handler
//}

func (chann *SubscribableChannel) Subscribe(handler stompmsg.Handler) error {
	return nil
}

func (chann *SubscribableChannel) Unsubscribe(handler stompmsg.Handler) error {
	return nil
}

func (chann *SubscribableChannel) Send(message stompmsg.Message[interface{}]) error {
	if message == nil {
		return wserror.IllegalArgument{Message: "Message must not be null"}
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

func (chann *SubscribableChannel) doConnectInternal(handler stompmsg.Handler) {
	//fmt.Printf("New client %s connected", handler.GetUserID())
	//chann.OutBoundChannels[handler.GetUserID()] = []socketmsg.Handler[[]byte]{handler}
}

func (chann *SubscribableChannel) doDisConnectInternal(handler stompmsg.Handler) {
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
