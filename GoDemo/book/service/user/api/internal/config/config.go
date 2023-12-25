package config
// 读取配置文件
import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)


type Config struct {
	// 原始配置
	rest.RestConf
	// 自定义配置
	Mysql struct{
		DataSource string
	}
	CacheRedis cache.CacheConf
	// jwt
	Auth struct{
		AccessSecret string
		AccessExpire int64
	}

}
