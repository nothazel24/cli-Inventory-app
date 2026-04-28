package errors

type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(msg string) error {
	return &AppError{Message: msg}
}
