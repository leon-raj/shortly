package apperr

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(code string, message string) *Error {
	return &Error{code, message}
}

// Shared app apperr, feels weird rn to put it all together, might come up with something better later
var (
	ErrNotFound      = New("NOT_FOUND", "url not found")
	ErrAlreadyExists = New("ALREADY_EXISTS", "the alias is already taken")
)
