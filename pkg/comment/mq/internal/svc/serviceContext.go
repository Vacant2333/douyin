package svc

import (
	"douyin/pkg/comment/mq/internal/config"
	"douyin/pkg/comment/rpc/usercomment"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserCommentRpc usercomment.UserComment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserCommentRpc: usercomment.NewUserComment(zrpc.MustNewClient(c.UserCommentRpc)),
	}
}
