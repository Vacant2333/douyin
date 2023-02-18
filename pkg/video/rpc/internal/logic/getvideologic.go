package logic

import (
	"context"

	"douyin/pkg/video/rpc/internal/svc"
	"douyin/pkg/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoLogic) GetVideo(in *video.GetVideoReq) (*video.GetVideoResp, error) {
	// todo: add your logic here and delete this line

	return &video.GetVideoResp{}, nil
}
