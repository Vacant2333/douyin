package listen

import (
	"context"
	"douyin/pkg/comment/mq/internal/config"
	kqMq "douyin/pkg/comment/mq/internal/mqs/kq"
	"douyin/pkg/comment/mq/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

//pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.UserCommentOptServiceConf, kqMq.NewUserCommentUpdateMq(ctx, svcContext)),
		//.....
	}

}
