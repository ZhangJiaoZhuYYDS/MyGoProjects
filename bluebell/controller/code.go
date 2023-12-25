// @Author zhangjiaozhu 2023/6/4 9:21:00
package controller

type ResCode int64

/*响应状态信息码*/
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken      ResCode = 1006
	CodeInvalidAuthFormat ResCode = 1007
	CodeNotLogin          ResCode = 1008
)

var codeMsg = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户已经存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码不正确",
	CodeServerBusy:        "服务器繁忙",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
}

func (c ResCode) Msg() string {
	s, ok := codeMsg[c]
	if !ok {
		s = codeMsg[CodeServerBusy]
	}
	return s
}
