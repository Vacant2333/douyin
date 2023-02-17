package svc

import (
	"douyin/pkg/comment/api/internal/config"
	"douyin/pkg/comment/api/internal/middleware"
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/userinfo-demo/rpc/userinfoclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserCommentRpc usercomment.UserComment
	UserInfoRpc    userinfoclient.Userinfo
	AuthJWT        rest.Middleware
	IsLogin        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserCommentRpc: usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
		UserInfoRpc:    userinfoclient.NewUserinfo(zrpc.MustNewClient(c.UserInfoRpc)),
		AuthJWT:        middleware.NewAuthJWTMiddleware().Handle,
		IsLogin:        middleware.NewIsLoginMiddleware().Handle,
	}
}
