package logic

import (
	"context"

	"douyin/pkg/follow/rpc/internal/svc"
	"douyin/pkg/follow/rpc/types/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *follow.FollowReq) (*follow.FollowResp, error) {
	// todo: add your logic here and delete this line

	return &follow.FollowResp{}, nil
}
