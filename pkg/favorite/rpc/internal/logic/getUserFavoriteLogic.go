package logic

import (
	"context"

	"douyin/pkg/favorite/rpc/internal/svc"
	"douyin/pkg/favorite/rpc/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteLogic {
	return &GetUserFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userFavoriteList-----------------------
func (l *GetUserFavoriteLogic) GetUserFavorite(in *userOptPb.GetUserFavoriteReq) (*userOptPb.GetUserFavoriteResp, error) {
	// todo: add your logic here and delete this line

	return &userOptPb.GetUserFavoriteResp{}, nil
}
