// @Author zhangjiaozhu 2023/3/4 20:23:00
package routers

import (
	"Buble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		//v1.POST("/login", service.LoginHandler)
		v1.POST("/register", controller.RegisterHandler)  // 注册
		v1.POST("/login", controller.LoginHandler)   // 登录
	}
	//v2 := r.Group("api/v2", middlewares.JWTAuthMiddleware()) // 添加token校验
	v2 := r.Group("api/v2") // 添加token校验
	{
		v2.GET("test", func(context *gin.Context) {
			context.JSON(200, "6666")
		})
		v2.GET("/community", controller.CommunityHandler) // 获取分类社区列表
		v2.GET("/community/:id", controller.CommunityHandlerDetail) // 获取分类社区列表

		v2.POST("/post",controller.CreatePostController) // 创建帖子

																	// 查询帖子详情
		v2.GET("/post2",controller.PostListsController)	// 分页查询帖子

		v2.POST("/vote",controller.VoteController) // 投票


	}
	//r.LoadHTMLGlob("./templates/")  // 加载html
	//r.Static("/static", "./static") // 加载静态文件

	return r
}
