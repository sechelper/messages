package messages

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type StatusCode int32

const (
	SuccessStatusCode             StatusCode = 20000
	BadRequestStatusCode                     = 40000
	ServerInternalErrorStatusCode            = 50000
	UnknownErrorStatusCode                   = 90000
)

type Message struct {
	Code StatusCode  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewMessage(code StatusCode, text string, body interface{}) *Message {
	return &Message{code, text, body}
}

func (msg *Message) GetMsg() string {
	return msg.Msg
}

func (msg *Message) StatusCode() StatusCode {
	return msg.Code
}

func (msg *Message) GetData() interface{} {
	return msg.Data
}

func Response(w http.ResponseWriter, data interface{}, err error) {
	var msg = Message{
		Code: SuccessStatusCode,
		Msg:  "success",
	}
	if err != nil {
		var errorMassage ErrorMassage
		switch {
		case errors.As(err, &errorMassage):
			msg.Code = err.(ErrorMassage).GetCode()
			msg.Msg = err.(ErrorMassage).GetMsg()
		}
	}

	msg.Data = data

	httpx.OkJson(w, msg)
}
