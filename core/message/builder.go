package message

type Builder[P interface{}, H interface{}] struct {
	Payload  P
	Accessor *Accessor
}

func NewBuilderWithPayload[P interface{}, H interface{}](payload P) *Builder[P, H] {
	return &Builder[P, H]{
		Payload:  payload,
		Accessor: NewAccessor(),
	}
}

func (b *Builder[P, T]) Build() Message[interface{}, interface{}] {
	return &GenericMessage[interface{}, interface{}]{
		Payload: b.Payload,
		Headers: b.Accessor.ToMessageHeaders(),
	}
}
