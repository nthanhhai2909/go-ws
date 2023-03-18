package header

type Headers struct {
	properties map[string]string
}

func EmptyHeader() *Headers {
	return &Headers{properties: make(map[string]string, 0)}
}

func WithProperties(properties map[string]string) *Headers {
	return &Headers{properties}
}

func (h *Headers) AddHeader(key string, value string) {
	h.properties[key] = value
}

func (h *Headers) GetHeader(key string) string {
	return h.properties[key]
}

func (h *Headers) Properties() map[string]string {
	return h.properties
}
