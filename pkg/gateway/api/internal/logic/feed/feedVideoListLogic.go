package feed

import (
	"context"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedVideoListLogic {
	return &FeedVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedVideoListLogic) FeedVideoList(req *types.FeedVideoListReq) (resp *types.FeedVideoListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
