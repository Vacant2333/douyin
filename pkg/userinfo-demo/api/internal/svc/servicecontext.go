package svc

import (
	"douyin/pkg/userinfo-demo/api/internal/config"
	"douyin/pkg/userinfo-demo/rpc/userinfoclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserinfoRpc userinfoclient.Userinfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserinfoRpc: userinfoclient.NewUserinfo(zrpc.MustNewClient(c.UserinfoRpc)),
	}
}
