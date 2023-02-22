package svc

import (
	"douyin/pkg/comment/usercomment"
	"douyin/pkg/favorite/useroptservice"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/gateway/internal/config"
	"douyin/pkg/gateway/internal/middleware"
	"douyin/pkg/message/usermessage"
	"douyin/pkg/user/userservice"
	"douyin/pkg/video/videoservice"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc         userservice.UserService
	UserCommentRpc  usercomment.UserComment
	UserFavoriteRpc useroptservice.UserOptService
	VideoRPC        videoservice.VideoService
	FollowRPC       followservice.FollowService
	MessageRpc      usermessage.UserMessage

	CommentOptMsgProducer  *kq.Pusher
	FavoriteOptMsgProducer *kq.Pusher
	FollowOptMsgProducer   *kq.Pusher

	AuthJWT rest.Middleware
	IsLogin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthJWT: middleware.NewAuthJWTMiddleware().Handle,
		IsLogin: middleware.NewIsLoginMiddleware().Handle,

		UserRpc:                userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		UserCommentRpc:         usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
		FollowRPC:              followservice.NewFollowService(zrpc.MustNewClient(c.FollowRPC)),
		FavoriteOptMsgProducer: kq.NewPusher(c.UserFavoriteOptServiceConf.Brokers, c.UserFavoriteOptServiceConf.Topic),
		MessageRpc:             usermessage.NewUserMessage(zrpc.MustNewClient(c.MessageRpc)),
		UserFavoriteRpc:        useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserFavoriteRpc)),
		VideoRPC:               videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRPC)),
		CommentOptMsgProducer:  kq.NewPusher(c.UserCommentOptServiceConf.Brokers, c.UserCommentOptServiceConf.Topic),
	}
}
