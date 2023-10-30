package message

import "fmt"

type ErrorMassage struct {
	code StatusCode
	text string
}

func NewErrorMessage(code StatusCode, text string) *ErrorMassage {
	return &ErrorMassage{
		code: code,
		text: text,
	}
}

func (err ErrorMassage) GetCode() StatusCode {
	return err.code
}

func (err ErrorMassage) GetMsg() string {
	return err.text
}

func (err ErrorMassage) Error() string {
	return fmt.Sprintf("code:%d，text:%s", err.code, err.text)
}
