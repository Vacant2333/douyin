package userOpt

import (
	"context"

	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowOptLogic {
	return &FollowOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowOptLogic) FollowOpt(req *types.FollowOptReq) (resp *types.FollowOptRes, err error) {
	// todo: add your logic here and delete this line

	return
}
