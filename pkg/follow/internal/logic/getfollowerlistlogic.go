package logic

import (
	"context"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *follow.GetFollowerListReq) (*follow.GetFollowerListResp, error) {
	// todo: add your logic here and delete this line

	return &follow.GetFollowerListResp{}, nil
}
