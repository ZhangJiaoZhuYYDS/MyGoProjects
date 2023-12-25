package main

import (
	"Buble/commons/snowflake"
	viperConfig "Buble/commons/viper"
	"Buble/dao/mysql"
	"Buble/routers"
	"fmt"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("conf", "c", "", "apiserver conf file path.")
)

func main() {
	// 解析命令行指令
	pflag.Parse()
	// 初始化配置文件
	if err := viperConfig.Init(*cfg); err != nil {
		panic(err)
	}
	// 初始化并获取连接池对象
	_ = mysql.GetDB()
	// 主程序中使用Init函数进行初始化雪花算法
	if err := snowflake.Init("2023-03-01", 1); err != nil {
		fmt.Printf("init sonwflake failed, err: %v\n", err)
		return
	}
	r := routers.SetupRouter()
	r.Run()
}
