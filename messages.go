package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	SuccessStatusCode             = 20000
	BadRequestStatusCode          = 40000
	ServerInternalErrorStatusCode = 50000
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

func (resp *Body) SetStatusCode(code int) {
	resp.Code = code
}

func (resp *Body) SetMessage(msg string) {
	resp.Msg = msg
}

func (resp *Body) SetData(data interface{}) {
	resp.Data = data
}

func NewResponseMessage(code int, msg string, data interface{}) *Body {
	return &Body{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewSuccessMessage(data interface{}) *Body {
	return NewResponseMessage(SuccessStatusCode, "success", data)
}

func NewBadRequestMessage(msg string) *Body {
	return NewResponseMessage(BadRequestStatusCode, msg, nil)
}

func NewServerInternalErrorMessage(msg string) *Body {
	return NewResponseMessage(ServerInternalErrorStatusCode, msg, nil)
}
