package svc

import (
	"douyin/common/model/followModel"
	"douyin/pkg/follow/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel followModel.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FollowModel: followModel.NewFollowModel(conn, c.CacheRedis),
	}
}
