package svc

import (
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/favorite/rpc/useroptservice"
	"douyin/pkg/gateway/api/internal/config"
	middleware2 "douyin/pkg/gateway/api/internal/middleware"
	"douyin/pkg/user/rpc/userservice"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc         userservice.UserService
	UserCommentRpc  usercomment.UserComment
	UserFavoriteRpc useroptservice.UserOptService

	CommentOptMsgProducer  *kq.Pusher
	FavoriteOptMsgProducer *kq.Pusher

	AuthJWT rest.Middleware
	IsLogin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthJWT: middleware2.NewAuthJWTMiddleware().Handle,
		IsLogin: middleware2.NewIsLoginMiddleware().Handle,

		UserRpc:               userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		CommentOptMsgProducer: kq.NewPusher(c.UserCommentOptServiceConf.Brokers, c.UserCommentOptServiceConf.Topic),
		UserCommentRpc:        usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
	}
}
