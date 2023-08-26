package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"login/login_api/internal/config"
	"login/login_rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
