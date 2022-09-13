package message

const (
	IndefiniteTimeout int = -1
)

type Channel interface {
	Send(message Message[interface{}], timeout int64) bool
}
