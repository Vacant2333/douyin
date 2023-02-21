package svc

import (
	"douyin/common/model/videoModel"
	"douyin/pkg/video/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel videoModel.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: videoModel.NewVideoModel(conn, c.CacheRedis),
	}
}
