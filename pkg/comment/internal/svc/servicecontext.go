package svc

import (
	"douyin/common/model/commentModel"
	"douyin/pkg/comment/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	UserCommentModel commentModel.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserCommentModel: commentModel.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
