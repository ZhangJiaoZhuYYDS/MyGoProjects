package middlewares

import (
	"TouPiaoSystem/handler"
	"github.com/gin-gonic/gin"
)

func CheckCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie存储的值
		myCookie, err := c.Cookie("login_cookie")
		if err != nil {
			handler.SendReponse(c, err, "未登录")
			c.Abort()
		}
		// 把cookie存储的值取出（用户id）设为上下文访问属性
		c.Set("UserId",myCookie)
		c.Next()
	}
}
