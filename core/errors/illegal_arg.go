package errors

type IllegalArgument struct {
	Message string
}

func (e IllegalArgument) Error() string { return e.Message }
