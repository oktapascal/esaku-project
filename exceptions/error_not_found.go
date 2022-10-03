package exceptions

type ErrorNotFound struct {
	Error string
}

func NewErrorNotFound(error string) ErrorNotFound {
	return ErrorNotFound{Error: error}
}
