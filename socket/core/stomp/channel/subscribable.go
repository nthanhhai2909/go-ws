package channel

import (
	"mem-ws/socket/core/stomp/broker"
	"mem-ws/socket/core/stomp/smsg"
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
