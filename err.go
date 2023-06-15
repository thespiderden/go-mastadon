package masta

type APIError struct {
	Code int
	err  error
}

func (e *APIError) Error() string {
	return e.err.Error()
}
