package channel

import "mem-ws/core/subprotocols/stomp/smsg"

type Outbound interface {
	Send(message smsg.IMessage) error
}
