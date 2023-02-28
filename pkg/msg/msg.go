package msg

const DEFAULT_MSG string = ""
const DEFAULT_URL string = ""
const DEFAULT_DATA string = ""

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回错误信息
func SendResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 返回错误信息
func Error(msg string) *Response {
	return SendResponse(1001, msg, nil)
}

// 返回正确信息
func Success(msg string, data interface{}) *Response {
	return SendResponse(200, msg, nil)
}
