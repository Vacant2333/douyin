package svc

import (
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/useroptservice"
	"douyin/pkg/follow/followservice"
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
	FavoriteRPC useroptservice.UserOptService
	FollowRPC   followservice.FollowService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		UserPRC:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		FavoriteRPC: useroptservice.NewUserOptService(zrpc.MustNewClient(c.FavoriteRpc)),
		MinioRPC:    minioclient.NewMinIOClient(zrpc.MustNewClient(c.MinIOClientRpc)),
		FollowRPC:   followservice.NewFollowService(zrpc.MustNewClient(c.FollowRPC)),
		VideoModel:  videoModel.NewVideoModel(conn, c.CacheRedis),
	}
}
