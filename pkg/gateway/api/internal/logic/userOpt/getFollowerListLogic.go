package userOpt

import (
	"context"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.FollowerListReq) (resp *types.FollowerListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
