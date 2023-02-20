package svc

import (
	"douyin/common/model/favoriteModel"
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserFavoriteModel favoriteModel.FavoriteModel
	UserVideoModel    videoModel.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserFavoriteModel: favoriteModel.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		UserVideoModel:    videoModel.NewVideoModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
	}
}
