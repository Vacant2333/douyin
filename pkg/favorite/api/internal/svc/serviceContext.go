package svc

import (
	"douyin/pkg/favorite/api/internal/config"
	"douyin/pkg/favorite/api/internal/middleware"
	"douyin/pkg/favorite/rpc/useroptservice"
	"douyin/pkg/user/rpc/userservice"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                 config.Config
	AuthJWT                rest.Middleware
	IsLogin                rest.Middleware
	UserFavoriteRpc        useroptservice.UserOptService
	UserRpc                userservice.UserService
	FavoriteOptMsgProducer *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		AuthJWT:                middleware.NewAuthJWTMiddleware().Handle,
		IsLogin:                middleware.NewIsLoginMiddleware().Handle,
		FavoriteOptMsgProducer: kq.NewPusher(c.UserFavoriteOptServiceConf.Brokers, c.UserFavoriteOptServiceConf.Topic),
		UserFavoriteRpc:        useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserFavoriteRpc)),
		UserRpc:                userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
