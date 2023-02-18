package logic

import (
	"context"

	"douyin/pkg/follow/rpc/internal/svc"
	"douyin/pkg/follow/rpc/types/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *follow.GetFollowListReq) (*follow.GetFollowListResp, error) {
	// todo: add your logic here and delete this line

	return &follow.GetFollowListResp{}, nil
}
