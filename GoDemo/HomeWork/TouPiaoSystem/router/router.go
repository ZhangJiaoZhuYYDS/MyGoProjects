package router

import (
	"TouPiaoSystem/handler/user"
	"TouPiaoSystem/router/middlewares"
	"TouPiaoSystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 使用自定义中间件
	g.Use(mw...)
	// 404
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "404:找不到页面")
	})

	//加载项目下所有的html
	g.LoadHTMLGlob("./resources/**/*")

	g.Static("/static", "./static/")

	// 获取当前文件执行路径
	//if ex, err := os.Executable(); err == nil {
	//	fmt.Println("66666666666", filepath.Dir(ex))
	//}



	// index
	g.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	//login
	g.POST("/login", user.Login)
	// 校验用户是否登录
	userCheck := g.Group("/userCheck/", middlewares.CheckCookie())
	{
		// 投票页面
		userCheck.GET("/form", func(context *gin.Context) {
			context.HTML(http.StatusOK, "form.html", nil)
		})
		// 发送投票请求
		userCheck.POST("/post", service.CreateVote)
		// 获取当前用户投票信息
		userCheck.GET("/vote/signal",service.UserVote)
		// 获取所有用户投票信息
		userCheck.GET("/vote/all",service.AllUserVote)
		// 获取总票数
		userCheck.POST("/vote/total",service.TotalVote)
		// 获取所有票类型
		userCheck.GET("/vote/votetype",service.VoteType)
	}
	return g
}
