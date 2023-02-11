package svc

import (
	"comment/rpc/internal/config"
	"comment/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
