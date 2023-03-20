package inbound

import (
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

// InboundManager TODO PROCESS CONCURRENCY HERE
// ENHANCE PERFORMANCE FOR UNSUBSCRIBE FLOW
type InboundManager struct {
	InboundMap map[string]IChannel
}

// TODO ALLOW USER CAN REGISTER HANDLER FOR DESTINATION SPECIFICLY
func (m *InboundManager) Send(destination string, message smsg.IMessage) error {
	if stringutils.IsBlank(destination) {
		return errors.IllegalArgument{Message: "Destination must not be null"}
	}

	if message == nil {
		return errors.IllegalArgument{Message: "Message must not be null"}
	}

	// TODO WORK WITH SEND METHOD

	return errors.IllegalArgument{Message: "Invalid Destination"}
}

func (m *InboundManager) Subscribe(msg smsg.IMessage, session session.ISession) error {
	if chann, ok := m.InboundMap[msg.GetMessageHeaders().Destination()]; ok {
		return chann.Subscribe(msg, session)
	}

	return errors.IllegalArgument{Message: "Invalid Destination"}
}

func (m *InboundManager) UnSubscribe(msg smsg.IMessage, session session.ISession) error {
	for _, inboud := range m.InboundMap {
		inboud.Unsubscribe(msg, session)
	}
	return nil
}
