// @Author zhangjiaozhu 2023/6/4 10:36:00
package models

// 定义注册请求参数
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

// 定义登录请求参数
type ParamLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
