package wscore

type Action string

const (
	SUBSCRIBE   Action = "SUBSCRIBE"
	UNSUBSCRIBE Action = "UNSUBSCRIBE"
	BROADCAST   Action = "BROADCAST"
)
