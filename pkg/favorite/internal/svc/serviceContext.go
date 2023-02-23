package svc

import (
	"douyin/common/model/favoriteModel"
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/internal/config"
	"douyin/pkg/video/videoservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	UserFavoriteModel favoriteModel.FavoriteModel
	UserVideoModel    videoModel.VideoModel
	VideoRPC          videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserFavoriteModel: favoriteModel.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		UserVideoModel:    videoModel.NewVideoModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
		VideoRPC:          videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRPC)),
	}
}
