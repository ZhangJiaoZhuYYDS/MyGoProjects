package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo2/mall/order/api/internal/config"
	"go-zero-demo2/mall/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
