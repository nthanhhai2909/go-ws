package channel

import (
	"fmt"
	"mem-ws/core/subprotocols/stomp/broker"
	"mem-ws/core/subprotocols/stomp/constans"
	"mem-ws/core/subprotocols/stomp/smsg"
)

type Subscribable struct {
	Broker broker.Broker
}

func (chann *Subscribable) Subscribe(handler smsg.Handler) error {
	return nil
}

func (chann *Subscribable) Unsubscribe(handler smsg.Handler) error {
	return nil
}

func (chann *Subscribable) Send(message smsg.IMessage) error {
	fmt.Println("Call tao nek: ", message)
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.CommandHeader))
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.StompDestinationHeader))
	return nil
}
