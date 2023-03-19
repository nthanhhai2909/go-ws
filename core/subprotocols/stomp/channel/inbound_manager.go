package channel

import (
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"mem-ws/core/errors"
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/session"
)

// InboundManager TODO PROCESS CONCURRENCY HERE
// ENHANCE PERFORMANCE FOR UNSUBSCRIBE FLOW
type InboundManager struct {
	InboundMap map[string]Inbound
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

func (m *InboundManager) Subscribe(destination string, subscribeId string, session session.ISession) error {

	if stringutils.IsBlank(destination) {
		return errors.IllegalArgument{Message: "Destination must not be null"}
	}

	if stringutils.IsBlank(subscribeId) {
		return errors.IllegalArgument{Message: "SubscribeId must not be null"}
	}

	if session == nil {
		return errors.IllegalArgument{Message: "Session must not be null"}
	}

	if chann, ok := m.InboundMap[destination]; ok {
		return chann.Subscribe(subscribeId, session)
	}

	return errors.IllegalArgument{Message: "Invalid Destination"}
}

func (m *InboundManager) UnSubscribe(subscribeId string, session session.ISession) error {
	if stringutils.IsBlank(subscribeId) {
		return errors.IllegalArgument{Message: "SubscribeId must not be null"}
	}

	if session == nil {
		return errors.IllegalArgument{Message: "Session must not be null"}
	}

	for _, inboud := range m.InboundMap {
		inboud.Unsubscribe(subscribeId, session)
	}
	return nil
}
