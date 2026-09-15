package errors

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(code string, message string) *Error {
	return &Error{code, message}
}
