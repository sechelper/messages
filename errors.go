package response

type Error struct {
	code StatusCode
	text string
}

func NewError(code StatusCode, text string) error {
	return &Error{code, text}
}

func (e *Error) Error() string {
	return e.text
}

func (e *Error) StatusCode() StatusCode {
	return e.code
}
