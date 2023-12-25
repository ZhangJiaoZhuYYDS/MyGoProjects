package main

import (
	"TouPiaoSystem/config"
	"TouPiaoSystem/model"
	"TouPiaoSystem/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("conf", "c", "", "apiserver conf file path.")
)

func main() {
	// 解析终端输入
	pflag.Parse()
	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 初始化数据库
	model.DB.Init()

	//user := model.UserModel{Username: "Jinzhu", Password: "18"}
	//model.DB.Self.AutoMigrate(&model.UserModel{})
	//result := model.DB.Self.Create(&user)
	//fmt.Println(result)

	// defer model.DB.Close()
	// 加载默认引擎
	g := gin.Default()
	// 全局中间件
	middlewares := []gin.HandlerFunc{}
	// 加载路由配置
	router.Load(g, middlewares...)
	err := g.Run()
	if err != nil {
		return
	}

}
