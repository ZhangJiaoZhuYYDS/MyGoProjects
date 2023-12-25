package main

import (
	"flag"
	"fmt"

	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/handler"
	"go-zero-demo/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	// 解析命令行
	flag.Parse()
	// go-zero读取配置信息
	var c config.Config
	conf.MustLoad(*configFile, &c)
	// go-zero加载服务器配置以及其他配置信息
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	// 实例化配置到上下文，相当于gin.contex
	ctx := svc.NewServiceContext(c)
	// 调用路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
