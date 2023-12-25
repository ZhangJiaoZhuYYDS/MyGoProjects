package svc

import (
	"go-zero-demo/greet/internal/config"
)
// 加载配置
type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
