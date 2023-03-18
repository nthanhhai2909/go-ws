package channel

import (
	"mem-ws/core/subprotocols/stomp/broker"
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
	return nil
}
