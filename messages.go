package response

import (
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type StatusCode int

const (
	SuccessStatusCode             StatusCode = 20000
	BadRequestStatusCode                     = 40000
	ServerInternalErrorStatusCode            = 50000
	UnknownErrorStatusCode                   = 90000
)

type Body struct {
	Code StatusCode  `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body = Body{
		Code: SuccessStatusCode,
		Msg:  "success",
		Data: resp,
	}
	var respError *Error
	if err != nil {
		if errors.As(err, &respError) {
			body.Code = respError.code
		} else {
			body.Code = UnknownErrorStatusCode
		}
		body.Msg = err.Error()
	}
	httpx.OkJson(w, body)
}

func (resp *Body) SetStatusCode(code StatusCode) {
	resp.Code = code
}

func (resp *Body) SetMessage(msg string) {
	resp.Msg = msg
}

func (resp *Body) SetData(data interface{}) {
	resp.Data = data
}
