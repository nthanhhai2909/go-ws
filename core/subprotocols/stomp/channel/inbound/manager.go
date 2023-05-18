package inbound

import (
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
	"sync"
)

// Manager TODO PROCESS CONCURRENCY HERE
// ENHANCE PERFORMANCE FOR UNSUBSCRIBE FLOW
type Manager struct {
	InboundMap sync.Map
}

// Send it can be send to specific user or group Subscribers
func (m *Manager) Send(destination string, message smsg.IMessage) error {
	if stringutils.IsBlank(destination) {
		return errors.IllegalArgument{Message: "Destination must not be null"}
	}

	if message == nil {
		return errors.IllegalArgument{Message: "Message must not be null"}
	}
	val, ok := m.InboundMap.Load(message.GetMessageHeaders().Destination())
	if !ok {
		return errors.IllegalArgument{Message: "Destination could not be found"}
	}
	val.(IChannel).Send(message)
	return nil
}

func (m *Manager) Subscribe(msg smsg.IMessage, session session.ISession) error {
	if val, ok := m.InboundMap.Load(msg.GetMessageHeaders().Destination()); ok {
		return val.(IChannel).Subscribe(msg, session)
	}

	return errors.IllegalArgument{Message: "Invalid Destination"}
}

func (m *Manager) UnSubscribe(msg smsg.IMessage, session session.ISession) error {
	m.InboundMap.Range(func(key, value any) bool {
		value.(IChannel).Unsubscribe(msg, session)
		return true
	})
	return nil
}

func (m *Manager) HandleConnectionClose(session session.ISession) error {
	m.InboundMap.Range(func(key, value any) bool {
		value.(IChannel).HandleConnectionClose(session)
		return true
	})
	return nil
}
