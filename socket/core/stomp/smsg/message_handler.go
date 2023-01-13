package smsg

type Handler interface {
	HandleMessage(msg Message[interface{}]) error
}
