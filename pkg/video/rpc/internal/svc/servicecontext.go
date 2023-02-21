package svc

import (
	"douyin/common/model/videoModel"
	"douyin/pkg/minio-client/minioclient"
	"douyin/pkg/user/userservice"
	"douyin/pkg/video/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel videoModel.VideoModel
	MinioRPC   minioclient.MinIOClient
	UserPRC    userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: videoModel.NewVideoModel(conn, c.CacheRedis),
	}
}
