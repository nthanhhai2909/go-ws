package core

type Subscriber[T interface{}] interface {
	HandlerMessage(msg Message[T]) error
}
