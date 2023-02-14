package svc

import (
	"douyin/pkg/comment/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
import "douyin/pkg/comment/rpc/model"

type ServiceContext struct {
	Config           config.Config
	UserCommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserCommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
