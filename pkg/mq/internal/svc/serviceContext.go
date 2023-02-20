package svc

import (
	"douyin/pkg/comment/usercomment"
	"douyin/pkg/favorite/useroptservice"
	"douyin/pkg/mq/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	UserCommentRpc  usercomment.UserComment
	UserFavoriteRpc useroptservice.UserOptService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserCommentRpc:  usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
		UserFavoriteRpc: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserFavoriteRpc)),
	}
}
