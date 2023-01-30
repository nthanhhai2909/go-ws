package socket

type MessageHandler interface {
	HandleMessage(message interface{}) (interface{}, error)
}
