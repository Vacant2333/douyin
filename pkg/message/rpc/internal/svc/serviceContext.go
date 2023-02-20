package svc

import (
	"douyin/common/model/messageModel"
	"douyin/pkg/message/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel messageModel.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MessageModel: messageModel.NewMessageModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
