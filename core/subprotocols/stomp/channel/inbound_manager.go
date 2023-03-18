package channel

import (
	"mem-ws/core/errors"
	"mem-ws/native/session"
)

type InboundManager struct {
	InboundMap map[string]Inbound
}

func (m *InboundManager) Subscribe(destination string, session session.ISession) error {
	if chann, ok := m.InboundMap[destination]; ok {
		return chann.Subscribe(session)
	}

	return errors.IllegalArgument{Message: "Invalid Destination"}
}

func (m *InboundManager) UnSubscribe(destination string, session session.ISession) error {
	if chann, ok := m.InboundMap[destination]; ok {
		return chann.Unsubscribe(session)
	}

	return errors.IllegalArgument{Message: "Invalid Destination"}
}
