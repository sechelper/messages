package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type StatusCode int

const (
	SuccessStatusCode             StatusCode = 20000
	BadRequestStatusCode                     = 40000
	ServerInternalErrorStatusCode            = 50000
)

type Body struct {
	Code StatusCode  `json:"code"`
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

func (resp *Body) SetStatusCode(code StatusCode) {
	resp.Code = code
}

func (resp *Body) SetMessage(msg string) {
	resp.Msg = msg
}

func (resp *Body) SetData(data interface{}) {
	resp.Data = data
}
