package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//// 使用BasicAuth中间件
//
//var secrets = gin.H{
//	"gavin": gin.H{"email": "sss@chen.com", "phone": "1821111222"},
//	"tom":   gin.H{"email": "sss22@chen.com", "phone": "1821111666"},
//}
//
//func main() {
//
//	r := gin.Default()
//
//	// 在组件中使用gin.BasicAuth()中间件
//	//gin.Accounts 是 map[string]string 的快捷方式
//
//	authorized := r.Group("/bb", gin.BasicAuth(gin.Accounts{
//		"gavin": "6789",
//		"tom":   "2345",
//	}))
//
//	authorized.GET("/aaa", func(c *gin.Context) {
//
//		// 获取user,他是在basicAuth中设定的
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		if secret, ok := secrets[user]; ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
//		}
//	})
//
//	r.Run(":3000")
//
//}
