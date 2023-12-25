// @Author zhangjiaozhu 2023/6/4 10:21:00
package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
)

// 注册
func SignUpHandler(c *gin.Context) {
	// 1 请求参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(p.Username) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2 业务处理
	logic.SignUp(p)
	// 3 返回响应
	ResponseSuccess(c, nil)
}

// 登录
func LoginHandler(c *gin.Context) {
	u := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&u); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(u.Username) == 0 || len(u.Password) == 0 {
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := logic.Login(u)
	if err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, "用户名或者密码错误")
		return
	}
	ResponseSuccess(c, "登陆成功")
}
