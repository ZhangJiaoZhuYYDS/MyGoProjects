package router

import (
	"RESTful__API/handler/sd"
	"RESTful__API/handler/user"
	"RESTful__API/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares. gin相关中间件

	// gin.Recovery（）:在处理某些请求时可能因为程序bug或者其他异常情况导致程序panic,这时候为
	//了不影响下一次请求的调用，需要通过gin.Recovery() 来恢复API服务器
	g.Use(gin.Recovery())
	// 强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	//浏览器跨域options设置
	g.Use(middleware.Options)
	//一些安全设置
	g.Use(middleware.Secure)

	g.Use(mw...)
	// 404 Handler. 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	// login请求
	g.POST("/login", user.Login)
	// 添加测试自定义错误的路由组
	//u := g.Group("/v1/user"){u.POST("", user.Create)}
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware()) // 路由中间件
	{
		u.POST("", user.Create)       //创建用户
		u.DELETE("/:id", user.Delete) //删除用户
		u.PUT("/:id", user.Update)    //更新用户
		u.GET("", user.List)          //用户列表
		u.GET("/:username", user.Get) //获取指定用户的详细信息
	}
	// The health check handlers 路由分组请求“/sd” 到不同方法请求
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
