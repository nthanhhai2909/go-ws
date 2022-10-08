package message

type MType string

const (
	Connect       MType = "CONNECT"
	ConnectACK    MType = "CONNECT_ACK"
	Subscribe     MType = "SUBSCRIBE"
	Unsubscribe   MType = "UNSUBSCRIBE"
	HeartBeat     MType = "HEARTBEAT"
	Disconnect    MType = "DISCONNECT"
	DisconnectAck MType = "DISCONNECT_ACK"
	Other         MType = "OTHER"
)
