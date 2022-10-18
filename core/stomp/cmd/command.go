package cmd

import "log"

type Command struct {
	Type string
}

var commandMap map[string]struct{}

const (
	Connect       string = "CONNECT"
	ConnectACK    string = "CONNECT_ACK"
	Subscribe     string = "SUBSCRIBE"
	Unsubscribe   string = "UNSUBSCRIBE"
	HeartBeat     string = "HEARTBEAT"
	Disconnect    string = "DISCONNECT"
	DisconnectAck string = "DISCONNECT_ACK"
	Other         string = "OTHER"
)

func init() {
	commandMap = make(map[string]struct{}, 0)
	commandMap["CONNECT"] = struct{}{}
	commandMap["CONNECT_ACK"] = struct{}{}
	commandMap["SUBSCRIBE"] = struct{}{}
	commandMap["UNSUBSCRIBE"] = struct{}{}
	commandMap["HEARTBEAT"] = struct{}{}
	commandMap["DISCONNECT"] = struct{}{}
	commandMap["DISCONNECT_ACK"] = struct{}{}
	commandMap["OTHER"] = struct{}{}
}

func ToCommand(command string) *Command {
	if _, found := commandMap[command]; !found {
		log.Panic("Invalid Command")
	}
	return &Command{Type: command}
}
