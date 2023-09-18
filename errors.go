package response

func New(code StatusCode, text string) error {
	return &errorString{code, text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	code StatusCode
	text string
}

func (e *errorString) Error() string {
	return e.text
}
