// @Author zhangjiaozhu 2023/3/4 20:37:00
package controller

import (
	"Buble/models"
	"Buble/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// RegisterHandler
//
//	@Description: 用户注册
//	@Author zhangjiaozhu 2023-03-05 12:21:30
//	@param c
func RegisterHandler(c *gin.Context) {
	// 1 获取参数 2 校验数据有效性
	var fo *models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有无，直接返回响应
		log.Println("register with invalid param", err)
		return
	}
	// 3 注册用户
	if err := service.Register(fo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error:注册用户失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
func LoginHandler(c *gin.Context) {
	// 1 获取参数  2 校验数据有效性
	var u *models.LoginForm
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回
		log.Println("请求参数有错误")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有错误",
		})
		return
	}
	// 业务处理登录
	atoken, err := service.Login(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "账号或密码错误",
		})
		return
	}
	//返回token
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": atoken,
	})
}
