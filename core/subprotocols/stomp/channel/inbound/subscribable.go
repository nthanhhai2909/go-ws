package inbound

import (
	"fmt"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/broker"
	"mem-ws/core/subprotocols/stomp/channel/inbound/subscriber"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

type Subscribable struct {
	Broker      broker.Broker
	Subscribers map[subscriber.Key]subscriber.Context
}

func (chann *Subscribable) Subscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	headers := msg.GetMessageHeaders()
	subscribeId := headers.ID()
	fmt.Println("ACK: ", headers.Ack())
	key := subscriber.Key{SubscribeID: subscribeId, SessionID: sessionId}
	if _, ok := chann.Subscribers[key]; ok {
		return errors.IllegalArgument{Message: "Session already exists!"}
	}
	chann.Subscribers[key] = subscriber.Context{Ack: headers.Ack(), Session: session}
	fmt.Println("Subscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Unsubscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	subscribeId := msg.GetMessageHeaders().ID()
	key := subscriber.Key{SubscribeID: subscribeId, SessionID: sessionId}
	if _, ok := chann.Subscribers[key]; !ok {
		return errors.IllegalArgument{Message: "Session does not exist!"}
	}
	delete(chann.Subscribers, key)
	fmt.Println("Unsubscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Send(message smsg.IMessage) error {
	fmt.Println(string(message.GetPayload()))
	return nil
}
