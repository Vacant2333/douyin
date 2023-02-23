package svc

import (
	"douyin/common/model/commentModel"
	"douyin/pkg/comment/internal/config"
	"douyin/pkg/video/videoservice"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UserCommentModel commentModel.CommentModel
	VideoRPC         videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserCommentModel: commentModel.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		VideoRPC:         videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRPC)),
	}
}
