package main

import (
	"github.com/lexkong/log"
	"vote_system/router"
	"vote_system/utils/mysql"
	"vote_system/utils/redis"
	"vote_system/utils/viper"
)

var err error
func main() {


	// 初始化文件配置
	err = viper.InitConfig()
	if err != nil {
		log.Info("解析文件配置失败")
	}
	// 初始化数据库
	err = mysql.InitDB()
	if err != nil {
		log.Info("获取数据池连接对象失败")
	}
	// 初始化redis
	err = redis.InitRedisDb()
	if err != nil {
		log.Info("redis 初始化失败")
	}
	r := router.Router()
	panic(r.Run(":8080"))

}
