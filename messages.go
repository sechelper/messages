package message

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
	var msg = new(Message)
	if err != nil {
		if errors.As(err, &ErrorMassage{}) {
			msg.Code = err.(ErrorMassage).GetCode()
			msg.Msg = err.(ErrorMassage).GetMsg()
		} else {
			msg.Msg = "错误请求"
		}
	}

	msg.Data = data

	httpx.OkJson(w, msg)
}
