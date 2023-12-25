// @Author zhangjiaozhu 2023/6/4 10:53:00
package main

import (
	"bluebell/common/snowflake"
	"bluebell/conf"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/router"
	"fmt"
)

func main() {
	//var confFile string
	//flag.StringVar(&confFile, "conf", "./conf/config.yaml", "配置文件")
	//flag.Parse()
	// 加载配置
	if err := conf.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(conf.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
