package stalker

const (
	SuccessStatusCode             = 20000
	BadRequestStatusCode          = 40000
	ServerInternalErrorStatusCode = 50000
)

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (resp *ResponseMessage) SetStatusCode(code int) {
	resp.Code = code
}

func (resp *ResponseMessage) SetMessage(msg string) {
	resp.Message = msg
}

func (resp *ResponseMessage) SetData(data interface{}) {
	resp.Data = data
}

func NewResponseMessage(code int, msg string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func NewSuccessMessage(data interface{}) *ResponseMessage {
	return NewResponseMessage(SuccessStatusCode, "success", data)
}

func NewBadRequestMessage() *ResponseMessage {
	return NewResponseMessage(BadRequestStatusCode, "Bad Request", nil)
}

func NewServerInternalErrorMessage() *ResponseMessage {
	return NewResponseMessage(ServerInternalErrorStatusCode, "Server Internal Error", nil)
}
