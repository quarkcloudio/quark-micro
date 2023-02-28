package msg

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeMap struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Ok                  = 200   // 执行成功
	ParameterError      = 10001 // 参数错误
	InternalServerError = 20201 // 内部服务调用异常
)

var CodeMaps = []*CodeMap{
	{
		Code: Ok,
		Msg:  "ok",
	},

	// 10001至19999，参数验证类状态码
	{
		Code: ParameterError,
		Msg:  "参数错误",
	},
	{
		Code: 10002,
		Msg:  "手机号格式错误",
	},
	{
		Code: 10003,
		Msg:  "邮箱格式错误",
	},
	{
		Code: 10004,
		Msg:  "短信验证码错误",
	},
	{
		Code: 10005,
		Msg:  "身份证格式错误",
	},
	{
		Code: 10006,
		Msg:  "参数不可为空",
	},
	{
		Code: 10007,
		Msg:  "账号或密码错误",
	},
	{
		Code: 10008,
		Msg:  "确认密码不一致",
	},
	{
		Code: 10009,
		Msg:  "图片大小或尺寸或文件类型不对",
	},
	{
		Code: 10010,
		Msg:  "图片上传失败",
	},
	// 20001至29999，项目业务类状态码
	{
		Code: 20001,
		Msg:  "用户不存在",
	},
	{
		Code: 20002,
		Msg:  "没查到数据",
	},
	{
		Code: 20003,
		Msg:  "手机短信验证码请求频繁",
	},
	{
		Code: 20004,
		Msg:  "用户已注销",
	},
	{
		Code: InternalServerError,
		Msg:  "内部服务调用异常",
	},
	// 30001至39999，数据库类状态码
	{
		Code: 30001,
		Msg:  "数据库操作失败",
	},
	{
		Code: 30002,
		Msg:  "数据库写入失败",
	},
	{
		Code: 30003,
		Msg:  "数据库更新失败",
	},
	{
		Code: 30004,
		Msg:  "数据库删除失败",
	},
	{
		Code: 30005,
		Msg:  "数据库查询失败",
	},
	// 40001至49999，会话类状态码
	{
		Code: 40001,
		Msg:  "账号或密码错误",
	},
	{
		Code: 40002,
		Msg:  "会话过期",
	},
	{
		Code: 40004,
		Msg:  "无效的访问令牌",
	},
	{
		Code: 40005,
		Msg:  "无效的邮箱验证码",
	},
	{
		// 例如用户授权给A客户端的code，只能在A客户端进行access_token的交换，用户在A客户端产生的refresh_token只能在A客户端进行刷新使用。
		Code: 40006,
		Msg:  "不匹配的客户端",
	},
	{
		Code: 40007,
		Msg:  "无效的Refresh Token",
	},
	{
		Code: 40008,
		Msg:  "错误的访问令牌（格式错误）",
	},
	{
		Code: 40009,
		Msg:  "Refresh Token格式错误",
	},
	{
		Code: 40010,
		Msg:  "Refresh Token解密失败",
	},
	{
		Code: 40011,
		Msg:  "无接口调用权限",
	},
	// 50001至59999，会权限类状态码
	{
		Code: 50001,
		Msg:  "未登录",
	},
	{
		Code: 50002,
		Msg:  "用户账号信息异常",
	},
	{
		Code: 50003,
		Msg:  "手机号码未认证",
	},
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
func Error(code int) *Response {
	if code == 200 {
		panic("错误状态码不能为:200")
	}
	return &Response{
		Code: code,
		Msg:  getCodeMsg(code),
		Data: nil,
	}
}

// 返回正确信息
func Success(data interface{}) *Response {
	code := 200
	return &Response{
		Code: code,
		Msg:  getCodeMsg(code),
		Data: data,
	}
}

// 获取当前code对应的信息
func getCodeMsg(code int) string {
	for _, v := range CodeMaps {
		if v.Code == code {
			return v.Msg
		}
	}

	return ""
}
