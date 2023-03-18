package core

type MessageHandler interface {
	HandleMessage(message interface{}) (interface{}, error)
}
