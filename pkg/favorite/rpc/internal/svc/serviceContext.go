package svc

import (
	"douyin/pkg/favorite/rpc/internal/config"
	"douyin/pkg/favorite/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserFavoriteModel model.FavoriteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserFavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
