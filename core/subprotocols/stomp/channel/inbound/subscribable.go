package inbound

import (
	"fmt"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/broker"
	"mem-ws/core/subprotocols/stomp/channel/inbound/subscriber"
	"mem-ws/core/subprotocols/stomp/cmd/server"
	"mem-ws/core/subprotocols/stomp/codec"
	"mem-ws/core/subprotocols/stomp/constants"
	"mem-ws/core/subprotocols/stomp/header"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
	"sync"
)

var encoder = codec.Encoder{}

type Subscribable struct {
	Broker      broker.Broker
	Subscribers sync.Map
}

func (chann *Subscribable) Subscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	headers := msg.GetMessageHeaders()
	subscribeId := headers.ID()
	key := subscriber.Key{SubscribeID: subscribeId, SessionID: sessionId}

	if _, ok := chann.Subscribers.Load(key); ok {
		return errors.IllegalArgument{Message: "Session already exists!"}
	}

	chann.Subscribers.Store(key, subscriber.Context{Ack: headers.Ack(), Session: session})
	return nil
}

func (chann *Subscribable) Unsubscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	subscribeId := msg.GetMessageHeaders().ID()
	key := subscriber.Key{SubscribeID: subscribeId, SessionID: sessionId}
	if _, ok := chann.Subscribers.Load(key); !ok {
		return errors.IllegalArgument{Message: "Session does not exist!"}
	}

	chann.Subscribers.Delete(key)
	return nil
}

func (chann *Subscribable) Send(message smsg.IMessage) error {
	chann.Subscribers.Range(func(key, value any) bool {
		ctx := value.(subscriber.Context)
		session := ctx.Session
		//ackMode := ctx.Ack
		subKey := key.(subscriber.Key)
		fmt.Println(subKey)
		session.SendMessage(encoder.Encode(smsg.NewGenericMessage(map[string]string{
			header.Command:           server.Message,
			header.StompContentType:  constants.TextPlain,
			header.StompSubscription: subKey.SubscribeID,
		}, message.GetPayload())))
		return true
	})
	return nil
}
