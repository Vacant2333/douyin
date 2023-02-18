package svc

import (
	"douyin/pkg/comment/api/internal/config"
	"douyin/pkg/comment/api/internal/middleware"
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/userinfo-demo/rpc/userinfoclient"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UserInfoRpc           userinfoclient.Userinfo
	UserCommentRpc        usercomment.UserComment
	CommentOptMsgProducer *kq.Pusher
	AuthJWT               rest.Middleware
	IsLogin               rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		UserInfoRpc:           userinfoclient.NewUserinfo(zrpc.MustNewClient(c.UserInfoRpc)),
		CommentOptMsgProducer: kq.NewPusher(c.UserCommentOptServiceConf.Brokers, c.UserCommentOptServiceConf.Topic),
		UserCommentRpc:        usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
		AuthJWT:               middleware.NewAuthJWTMiddleware().Handle,
		IsLogin:               middleware.NewIsLoginMiddleware().Handle,
	}
}
