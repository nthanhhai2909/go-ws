package socket

type Action string

const (
	SUBSCRIBE    Action = "SUBSCRIBE"
	UNSUBSCRIBE  Action = "UNSUBSCRIBE"
	BROADCAST    Action = "BROADCAST"
	SEND_TO_USER Action = "SEND_TO_USER"
)
