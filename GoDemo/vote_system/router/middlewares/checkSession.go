// @Author zhangjiaozhu 2023/3/18 23:47:00
package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)
// 校验是否存在session
func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session :=sessions.Default(c)
		// 通过 session.Get 读取 session 值
		userid := session.Get("UserId")
		if userid == nil {
			c.JSON(http.StatusOK,gin.H{
				"code":101,
				"msg":"未登录",
				"id":userid,
			})
			c.Abort()
			// 跳转到登录界面。。。。
			return
		}
		c.Set("user_id",userid)
		c.Next()
	}
}
