// @Author zhangjiaozhu 2023/3/5 12:39:00
package models

type Code int64

const (
	CodeSuccess         Code = 1000
	CodeInvalidParams   Code = 1001
	CodeUserExist       Code = 1002
	CodeUserNotExist    Code = 1003
	CodeInvalidPassword Code = 1004
	CodeServerBusy      Code = 1005

	CodeInvalidToken      Code = 1006
	CodeInvalidAuthFormat Code = 1007
	CodeNotLogin          Code = 1008
)

var msg = map[Code]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeInvalidToken:      "无效token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "用户未登录",
}

func (c Code) Msg() string {
	msg, ok := msg[c]
	if ok {
		return msg
	}
	return string(msg[CodeServerBusy])
}
