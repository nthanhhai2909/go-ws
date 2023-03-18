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
	Subscribers map[string]session.ISession
}

func (chann *Subscribable) Subscribe(session session.ISession) error {
	sessionId := session.GetID()
	if _, ok := chann.Subscribers[sessionId]; ok {
		return errors.IllegalArgument{Message: "Session already exists!"}
	}
	chann.Subscribers[sessionId] = session
	fmt.Println("Subscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Unsubscribe(session session.ISession) error {
	sessionId := session.GetID()
	if _, ok := chann.Subscribers[sessionId]; !ok {
		return errors.IllegalArgument{Message: "Session does not exist!"}
	}
	delete(chann.Subscribers, sessionId)
	fmt.Println("Unsubscribe: ", chann.Subscribers)
	return nil
}

func (chann *Subscribable) Send(message smsg.IMessage) error {
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.CommandHeader))
	fmt.Println(message.GetMessageHeaders().GetHeader(constans.StompDestinationHeader))
	return nil
}
