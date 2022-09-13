package message

type SendingOperations interface {
	Send(destination string, message Message[interface{}]) error
}
