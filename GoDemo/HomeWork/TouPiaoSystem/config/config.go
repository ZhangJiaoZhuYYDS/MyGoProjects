package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

// 初始化配置文件
func Init(cfg string) error {
	config := Config{Name: cfg}
	if err := config.initConfig(); err != nil {
		return err
	}
	// 初始化日志包
	//conf.initLog()
	// 监控配置文件并热加载程序
	//conf.watchConfig()
	return nil
}

// 加载配置文件
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf")   // 没有指定配置文件，则解析默认的配置文件所在路径
		viper.SetConfigName("config") // 设置文件名
	}
	viper.SetConfigType("yaml") //设置配置文件的格式为yaml
	// 通过指定配置文件可以很方便地连接不同的环境（开发环境、测试环境）并加载不同的配置，方便开发和测试。
	viper.AutomaticEnv() // 读取匹配的环境变量
	//viper.SetEnvPrefix("APISERVER") //读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil { // viper 解析配置文件
		fmt.Println("viper解析配置文件失败", err)
		return err
	}
	return nil
}
