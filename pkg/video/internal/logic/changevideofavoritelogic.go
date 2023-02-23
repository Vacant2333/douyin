package logic

import (
	"context"
	"douyin/pkg/logger"

	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeVideoFavoriteLogic {
	return &ChangeVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeVideoFavoriteLogic) ChangeVideoFavorite(in *video.ChangeVideoFavoriteReq) (*video.ChangeVideoFavoriteResp, error) {
	err := l.svcCtx.VideoModel.UpdateCount(l.ctx, in.VideoId, "like_count", in.ActionType)
	if err != nil {
		logger.Fatalf("ChangeVideoFavoriteCount failed %s", err.Error())
		return nil, err
	}
	return &video.ChangeVideoFavoriteResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
