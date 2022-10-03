package exceptions

type ErrorBadRequest struct {
	Error string
}

func NewErrorBadRequest(error string) ErrorBadRequest {
	return ErrorBadRequest{Error: error}
}
