package errors

type MessageConversion struct {
	Message string
}

func (e MessageConversion) Error() string { return e.Message }
