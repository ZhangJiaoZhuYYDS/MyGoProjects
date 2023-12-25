package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	type UserInfo struct {
		Name string		`json:"name"`
		Age int 		`json:"age"`
	}

	// 创建一个默认路由引擎  相当于对http库里面的listenAndServer以及handleFunc进行封装。用默认路由实现
	engine := gin.Default()
	// 加载静态文件
	engine.Static("/xxx","./static")
	// 找模板文件
	engine.LoadHTMLGlob("static/*")
	// gin中添加自定义函数模板(转换为html格式)
	engine.SetFuncMap(template.FuncMap{
		"safe" : func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	// get请求方式  "/hello":请求路径  当客户端以get请求时，会执行函数
	engine.GET("/hello", func(context *gin.Context) {
		engine.LoadHTMLGlob("static/*")  // 模板解析
		context.HTML(http.StatusOK,"hello.tmpl",gin.H{
			"title" :"<a href='http://www.'>垃圾</a>", // 渲染模板
		})
		context.JSON(200,gin.H{
			"message":"get请求",
		})
		fmt.Println("666")
	})
	engine.POST("/hello", func(context *gin.Context) {
		if ex , err :=os.Executable();err != nil {
			fmt.Println(filepath.Dir(ex))
		}
		fmt.Println("./")
		context.JSON(200,gin.H{
			"message":"post请求",
		})
		fmt.Println("666")
	})
	engine.PUT("/hello", func(context *gin.Context) {
		var Msg struct{
			Name string `json:"user"`
			Message string
		}
		Msg.Name = "王"
		Msg.Message="66"
		context.JSON(http.StatusOK,Msg)
		context.JSON(200,gin.H{
			"message":"put请求",
		})
		fmt.Println("666")
	})
	engine.DELETE("/hello", func(context *gin.Context) {
		var user UserInfo

		err :=context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"error" : err.Error(),
			})
		}else {
			fmt.Println(user)
		context.JSON(200,gin.H{
			"message":"del请求",
		})
		fmt.Println("666")
	}
	})
	group := engine.Group("v1")
	{
		group.GET("/get", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message" : 12 ,
			})
		})
		group.GET("/get1", func(context *gin.Context) {
			context.HTML(http.StatusOK,"demo.html",nil)
			context.JSON(http.StatusOK,gin.H{
				"message" : true,
			})
		})
	}

	// 启动服务，默认8080端口
	engine.Run(":9090")
}
