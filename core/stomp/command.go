package stomp

type Command string

const (
	Connect       Command = "CONNECT"
	ConnectACK    Command = "CONNECT_ACK"
	Subscribe     Command = "SUBSCRIBE"
	Unsubscribe   Command = "UNSUBSCRIBE"
	HeartBeat     Command = "HEARTBEAT"
	Disconnect    Command = "DISCONNECT"
	DisconnectAck Command = "DISCONNECT_ACK"
	Other         Command = "OTHER"
)
