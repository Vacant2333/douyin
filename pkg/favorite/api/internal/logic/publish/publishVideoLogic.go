package publish

import (
	"context"

	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PubVideoReq) (resp *types.PubVideoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
