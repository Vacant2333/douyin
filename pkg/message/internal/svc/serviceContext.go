package svc

import (
	"douyin/common/model/followModel"
	"douyin/common/model/messageModel"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/message/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel messageModel.MessageModel
	FollowModel  followModel.FollowModel
	FollowRPC    followservice.FollowService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: messageModel.NewMessageModel(conn, c.CacheConf),
		FollowModel:  followModel.NewFollowModel(conn, c.CacheConf),
		FollowRPC:    followservice.NewFollowService(zrpc.MustNewClient(c.FollowRPC)),
	}
}
