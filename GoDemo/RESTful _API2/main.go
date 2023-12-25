package main

import (
	"RESTful__API/config"
	"RESTful__API/model"
	"RESTful__API/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("conf", "c", "", "apiserver conf file path.")
)

func main() {

	pflag.Parse()

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db

	model.DB.Init()
	defer model.DB.Close()

	//set gin mode
	gin.SetMode(viper.GetString("runmode"))
	// 1 加载gin引擎
	g := gin.New()

	// gin 自定义引入全局相关中间件
	middlewares := []gin.HandlerFunc{}

	// 路由
	router.Load(
		//  cores
		g,
		//中间件
		middlewares...,
	)
	// 服务状态自检
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由请求没有响应，或者花费太长时间去请求,已自动关闭服务器", err)
		}
		log.Info("路由成功")
	}()
	//

	fmt.Printf("start to listening the incoming requests on http address: %s ,", viper.GetString("addr"))
	fmt.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// 服务器状态自检
/*有时候API进程起来不代表API服务器正常，笔者曾经就遇到过这种问题：API进程存在，但是服务器却不
能对外提供服务。因此在启动API服务器时，如果能够最后做一个自检会更好些。
在启动HTTP端口前go一个pingServer协程，启动HTTP端口后，该协程不断地ping/sd/health路径，
如果失败次数超过一定次数，则终止HTTP服务器进程。通过自检可以最大程
度地保证启动后的APl服务器处于健康状态。
*/
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		//get请求访问“/health”
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("正在自检，1s后再次尝试")
		time.Sleep(time.Second)
	}
	return errors.New("连接不上路由")
}
