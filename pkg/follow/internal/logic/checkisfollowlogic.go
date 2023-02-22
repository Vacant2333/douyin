package logic

import (
	"context"

	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFollowLogic {
	return &CheckIsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsFollowLogic) CheckIsFollow(in *follow.CheckIsFollowReq) (*follow.CheckIsFollowResp, error) {
	isFollow, err := l.svcCtx.FollowModel.CheckIsFollow(l.ctx, in.UserId, in.FunId)
	if err != nil {
		logx.Errorf("query is follow failed: %v", err.Error())
		return nil, err
	}
	return &follow.CheckIsFollowResp{
		StatusCode: 0,
		StatusMsg:  "success",
		IsFollow:   isFollow,
	}, nil
}
