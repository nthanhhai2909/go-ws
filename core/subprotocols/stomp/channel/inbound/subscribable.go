package inbound

import (
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

// Subscribable
// Subscribers structure: <key>:<value> <SessionID: string>:[]<Ack: string, session: Session, SubscribeID: string>
type Subscribable struct {
	Broker      broker.Broker
	Subscribers *sync.Map
}

func (chann *Subscribable) Subscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	headers := msg.GetMessageHeaders()
	subscribeId := headers.ID()
	val, ok := chann.Subscribers.Load(sessionId)
	if ok {
		ctx := val.(subscriber.Context)
		ctx.Subscription.Store(subscribeId, nil)
	} else {
		Subscription := &sync.Map{}
		Subscription.Store(subscribeId, nil)
		chann.Subscribers.Store(sessionId, subscriber.Context{Ack: headers.Ack(), Session: session, Subscription: Subscription})
	}
	return nil
}

func (chann *Subscribable) Unsubscribe(msg smsg.IMessage, session session.ISession) error {
	sessionId := session.GetID()
	subscribeId := msg.GetMessageHeaders().ID()
	val, ok := chann.Subscribers.Load(sessionId)

	if !ok {
		return errors.IllegalArgument{Message: "Session does not exist!"}
	}

	ctx := val.(subscriber.Context)
	_, subOK := ctx.Subscription.Load(subscribeId)
	if !subOK {
		return errors.IllegalArgument{Message: "Subscriber ID does not exist!"}
	}
	ctx.Subscription.Delete(subscribeId)
	return nil
}

// Send TODO HGA PROCESS ACK
func (chann *Subscribable) Send(message smsg.IMessage) error {
	chann.Subscribers.Range(func(key, value any) bool {
		ctx := value.(subscriber.Context)
		ctx.Subscription.Range(func(key, value any) bool {
			ctx.Session.SendMessage(encoder.Encode(smsg.NewGenericMessage(map[string]string{
				header.Command:           server.Message,
				header.StompContentType:  constants.TextPlain,
				header.StompSubscription: key.(string),
			}, message.GetPayload())))
			return true
		})
		return true
	})
	return nil
}
