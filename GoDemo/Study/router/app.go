// @Author zhangjiaozhu 2023/3/11 22:34:00
package router

import (
	_ "Study/docs"
	"Study/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
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
	//访问swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 问题相关
	r.GET("/ping", service.GetProblemList)  //根据条件查询相关问题列表
	r.GET("/problem-detail",service.GetProblemDetail)

	// 用户
	r.GET("/user-detail",service.GetUserDetail)
	r.POST("/login",service.Login)
	r.POST("/send-code",service.SendCode)
	r.POST("/register",service.Register)

	// 排行榜
	r.GET("/rank-list",service.GetRankList)

	// 提交记录
	r.GET("/submit-list",service.GetSubmitList)
	return r
}
