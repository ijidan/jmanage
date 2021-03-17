package pkg

//错误代码
const (
	SUCCESS       = 0
	ERROR         = 500
	InvalidParams = 400

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
)

//错误信息映射
var CodeMessageMap = map[int]string {
	SUCCESS :      "ok",
	ERROR :        "fail",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

//获取错误信息
func GetMessage(code int) string {
	msg, ok := CodeMessageMap[code]
	if ok {
		return msg
	}
	return CodeMessageMap[ERROR]
}