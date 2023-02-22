package svc

import (
	"douyin/common/model/followModel"
	"douyin/pkg/follow/internal/config"
	"douyin/pkg/user/userservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel followModel.FollowModel
	UserPRC     userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FollowModel: followModel.NewFollowModel(conn, c.CacheRedis),
		UserPRC:     userservice.NewUserService(zrpc.MustNewClient(c.UserRPC)),
	}
}
