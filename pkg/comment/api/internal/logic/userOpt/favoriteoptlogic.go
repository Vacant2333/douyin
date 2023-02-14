package userOpt

import (
	"context"

	"douyin/pkg/comment/api/internal/svc"
	"douyin/pkg/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteOptLogic {
	return &FavoriteOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteOptLogic) FavoriteOpt(req *types.FavoriteOptReq) (resp *types.FavoriteOptRes, err error) {
	// todo: add your logic here and delete this line

	return
}
