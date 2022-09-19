package simp

type Type string

const (
	Connect       Type = "CONNECT"
	ConnectACK    Type = "CONNECT_ACK"
	Subscribe     Type = "SUBSCRIBE"
	Unsubscribe   Type = "UNSUBSCRIBE"
	HeartBeat     Type = "HEARTBEAT"
	Disconnect    Type = "DISCONNECT"
	DisconnectAck Type = "DISCONNECT_ACK"
)
