package svc

import (
	"douyin/pkg/favorite/api/internal/config"
	"douyin/pkg/favorite/api/internal/middleware"
	"douyin/pkg/favorite/rpc/useroptservice"
	"douyin/pkg/userinfo-demo/rpc/userinfoclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	AuthJWT         rest.Middleware
	IsLogin         rest.Middleware
	UserFavoriteRpc useroptservice.UserOptService
	UserInfoRpc     userinfoclient.Userinfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthJWT:         middleware.NewAuthJWTMiddleware().Handle,
		IsLogin:         middleware.NewIsLoginMiddleware().Handle,
		UserFavoriteRpc: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserFavoriteRpc)),
		UserInfoRpc:     userinfoclient.NewUserinfo(zrpc.MustNewClient(c.UserInfoRpc)),
	}
}
