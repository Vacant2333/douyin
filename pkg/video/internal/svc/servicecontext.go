package svc

import (
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/useroptservice"
	"douyin/pkg/minio-client/minioclient"
	"douyin/pkg/user/userservice"
	"douyin/pkg/video/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	VideoModel  videoModel.VideoModel
	MinioRPC    minioclient.MinIOClient
	UserPRC     userservice.UserService
	FavoritePRC useroptservice.UserOptService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		UserPRC:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		FavoritePRC: useroptservice.NewUserOptService(zrpc.MustNewClient(c.FavoriteRpc)),
		MinioRPC:    minioclient.NewMinIOClient(zrpc.MustNewClient(c.MinIOClientRpc)),
		VideoModel:  videoModel.NewVideoModel(conn, c.CacheRedis),
	}
}
