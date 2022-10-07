package message

const (
	Destination = "h-destination"
)

type Headers[K string, V interface{}] struct {
	headers map[K]V
}

func (h *Headers[K, V]) SetDestination(destination V) {
	h.headers[Destination] = destination
}
