package svc

import (
	"douyin/common/model/followModel"
	"douyin/common/model/userModel"
	"douyin/pkg/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   userModel.UserModel
	FollowModel followModel.FollowModel
	RedisCache  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisCache:  c.RedisCacheConf.NewRedis(),
		UserModel:   userModel.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		FollowModel: followModel.NewFollowModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
