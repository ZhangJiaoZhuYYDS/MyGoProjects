// @Author zhangjiaozhu 2023/6/4 10:17:00
package router

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	// 使用中间件
	r.Use()
	v1 := r.Group("/api/v1")
	// 注册路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录路由
	v1.POST("/login", controller.LoginHandler)
	// 刷新token
	v1.GET("/refresh_token", controller.)

	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)

		v1.GET("/post2", controller.PostList2Handler)

		v1.POST("/vote", controller.VoteHandler)

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment", controller.CommentListHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
