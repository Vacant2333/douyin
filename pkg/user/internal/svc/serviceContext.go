package svc

import (
	"douyin/common/model/favoriteModel"
	"douyin/common/model/followModel"
	"douyin/common/model/userModel"
	"douyin/pkg/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     userModel.UserModel
	FollowModel   followModel.FollowModel
	FavoriteModel favoriteModel.FavoriteModel
	RedisCache    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		RedisCache:    c.RedisCacheConf.NewRedis(),
		UserModel:     userModel.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		FavoriteModel: favoriteModel.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		FollowModel:   followModel.NewFollowModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
