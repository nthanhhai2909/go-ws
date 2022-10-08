package message

// Accessor is Headers wrapper class to help us normalize header datas
type Accessor struct {
	headers Headers[interface{}]
}

// NewAccessor TODO HGA WILL ADAPT LATER
func NewAccessor() *Accessor {
	return &Accessor{}
}

func (a *Accessor) ToMessageHeaders() *Headers[interface{}] {
	return nil
}

func (a *Accessor) SetContentType(contentType string) {
	a.headers.SetContentType(contentType)
}

func (a *Accessor) SetDestination(destination string) {
	a.headers.SetDestination(destination)
}
