package subprotocol

import (
	"mem-ws/core/subprotocols/stomp/smsg"
	"mem-ws/native/message"
	"mem-ws/native/session"
)

// ISubProtocolHandler is specifications that all sub-protocol such as Stomp have to implement
type ISubProtocolHandler interface {
	SupportProtocols() []string

	HandleMessageFromClient(session session.ISession, message message.IMessage)

	SendMessageToClient(session session.ISession, message smsg.IMessage)
}
