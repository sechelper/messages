package messages

import (
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

func Response(w http.ResponseWriter, data interface{}, err error) {
	var msg = Message{
		Code: SuccessStatusCode,
		Msg:  "成功",
		Data: data,
	}

	if err != nil {
		if e, ok := err.(*ErrorMassage); ok {
			msg.Code = e.GetCode()
			msg.Msg = e.GetMsg()
		} else {
			msg = Message{
				Code: ServerInternalErrorStatusCode,
				Msg:  "服务器内部错误，请稍候重试！",
			}
		}
	}

	msg.Data = data

	httpx.OkJson(w, msg)
}
