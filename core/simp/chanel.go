package simp

type Channel[T interface{}] interface {
	Send(message Message[T], timeout int64) error
}
