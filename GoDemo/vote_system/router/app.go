// @Author zhangjiaozhu 2023/3/18 13:35:00
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"log"
	_ "vote_system/docs"
	"vote_system/models"
	"vote_system/router/middlewares"
	"vote_system/service"
	"vote_system/utils/mysql"
)
// @title 标题
// @version 1.0
// @description 描述信息
// @termsOfService http://swagger.io/terms/
// @contact.name 联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func Router() *gin.Engine {
	r := gin.Default()
	//设置跨域请求
	config := cors.DefaultConfig()
	//允许所有跨域请求
	//config.AllowOrigins = []string{"http://localhost:63342","http://myhost:8090"}
	config.AllowOrigins = []string{"http://myhost:8090"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET, POST, PUT, DELETE, OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	r.Use(cors.New(config))
	// 初始化基于 redis 的存储引擎
	// size:redis 最大的空闲连接数
	//network: 数通信协议tcp或者udp
	//address:redis 地址, 格式，host:port
	//password:redis密码
	//最后一个参数：session 加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("SecretKey"))
	// 使用redis存储session
	//store := redis.NewStore([]byte("SecretKey"))
	r.Use(sessions.Sessions("sessionid", store))
	// 设置静态资源
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static","./static/*")
	//访问swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 测试跨域获取session
	r.GET("/test", func(c *gin.Context) {
		if  sessions.Default(c).Get("hello") != "world" {
			sessions.Default(c).Set("hello","world")
			sessions.Default(c).Save()
		}
		c.JSON(200,gin.H{
			"msg":66,
			"sessin":sessions.Default(c).Get("hello"),
		})
	})
	// 用户登录
	r.POST("/login",service.Login)
	r.POST("/logins",service.Login2)
	// 登录页面
	r.GET("/loginIndex", func(c *gin.Context) {
		c.HTML(200,"login.html",nil)
	})
	// 注册
	r.POST("/register",service.Register)
	// 发送邮箱验证码
	r.POST("/sendEmailCode", func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Next()
	},service.SendEmailCode)
	// 发送邮箱验证码
	r.POST("/sendemailcode",service.SendEmailCode2)
	// 原来投票路由组，不知道为啥就是无法访问，其他路径都行
	votesGroup := r.Group("/votes",middlewares.CheckSession())
	{	// 测试
		votesGroup.GET("/test", func(context *gin.Context) {
			session := sessions.Default(context)
			seId := session.Get("UserId")
			value,exist := context.Get("UserId")
			if !exist {
				log.Println(">>>>>>>>>>>>>>")
			}
			context.JSON(200,gin.H{
				"session":seId,
				"c":value,
				"msg":666,
			})
		})
		// 用户退出
		votesGroup.GET("/logout",service.Logout)
		// 用户投票
		votesGroup.POST("/doVote", service.DoVote)
		// 获取单个类型的所有用户投票信息
		votesGroup.POST("/getVotesByType",service.GetVotesByType)
		// 获取当前用户的某个投票类型的投票的信息
		votesGroup.POST("/getVotesoptByType",service.GetVotesoptByType)
		// 获取所有类型的投票结果
		votesGroup.GET("/getVotes",service.GetVotes)
		// 创建投票
		votesGroup.POST("/createVote",service.CreateVote)
	}
	// 投票路由组的替代，可以通过
	gr :=r.Group("/group",middlewares.CheckSession())
	{
		// 测试
		gr.GET("/api", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"msg":666,
			})
		})
		// 测试获取当前登录用户的信息
		gr.GET("/test", func(context *gin.Context) {
			session := sessions.Default(context)
			seId := session.Get("UserId")
			if seId == "" {
				context.JSON(200,gin.H{
					"code":101,
					"msg":"未登录",
				})
				return
			}
			var s string
			mysql.DB.Model(&models.User{}).Select("username").Where("user_identity = ?",seId).First(&s)
			context.JSON(200,gin.H{
				"session":seId,
				"c":seId,
				"msg":s,
			})
		})
		// 用户退出
		gr.GET("/logout",service.Logout)
		// 用户投票
		gr.POST("/doVote", service.DoVote)
		// 获取单个类型的所有用户投票信息
		gr.POST("/getVotesByType",service.GetVotesByType)
		// 获取当前用户的某个投票类型的投票的信息
		gr.POST("/getVotesoptByType",service.GetVotesoptByType)
		// 获取所有类型的投票结果
		gr.GET("/getVotes",service.GetVotes)
		// 创建投票
		gr.POST("/createVote",service.CreateVote)
	}
	return r
}
