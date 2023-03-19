package channel

import (
	"fmt"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/broker"
	"mem-ws/core/subprotocols/stomp/constans"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

type Subscribable struct {
	Broker      broker.Broker
	Subscribers map[SubscriberKey]session.ISession
}

type SubscriberKey struct {
	SessionID   string
	SubscribeID string
}

func (chann *Subscribable) Subscribe(subscribeId string, session session.ISession) error {
	sessionId := session.GetID()
	key := SubscriberKey{SubscribeID: subscribeId, SessionID: sessionId}
	if _, ok := chann.Subscribers[key]; ok {
		return errors.IllegalArgument{Message: "Session already exists!"}
	}
	chann.Subscribers[key] = session
	fmt.Println("Subscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Unsubscribe(subscribeId string, session session.ISession) error {
	sessionId := session.GetID()
	key := SubscriberKey{SubscribeID: subscribeId, SessionID: sessionId}
	if _, ok := chann.Subscribers[key]; !ok {
		return errors.IllegalArgument{Message: "Session does not exist!"}
	}
	delete(chann.Subscribers, key)
	fmt.Println("Unsubscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Send(message smsg.IMessage) error {
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.CommandHeader))
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.StompDestinationHeader))
	fmt.Println(string(message.GetPayload()))
	return nil
}
