package message

type Headers[K string, V interface{}] struct {
	Headers map[K]V
}
