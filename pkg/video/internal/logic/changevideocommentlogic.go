package logic

import (
	"context"
	"douyin/pkg/logger"

	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeVideoCommentLogic {
	return &ChangeVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeVideoCommentLogic) ChangeVideoComment(in *video.ChangeVideoCommentReq) (*video.ChangeVideoCommentResp, error) {
	err := l.svcCtx.VideoModel.UpdateCount(l.ctx, in.VideoId, "comment_count", in.ActionType)
	if err != nil {
		logger.Fatalf("ChangeVideoCommentCount failed %s", err.Error())
		return nil, err
	}
	return &video.ChangeVideoCommentResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
