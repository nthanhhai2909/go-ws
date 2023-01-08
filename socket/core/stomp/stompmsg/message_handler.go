package stompmsg

type Handler interface {
	HandleMessage(msg Message[interface{}]) error
}
