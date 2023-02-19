package svc

import (
	"douyin/pkg/user/api/internal/config"
	"douyin/pkg/user/api/internal/middleware"
	"douyin/pkg/user/rpc/userservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
	AuthJWT rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthJWT: middleware.NewAuthJWTMiddleware().Handle,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
