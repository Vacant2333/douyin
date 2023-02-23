package logic

import (
	"context"
	"douyin/pkg/logger"

	"douyin/pkg/favorite/internal/svc"
	"douyin/pkg/favorite/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFavoriteLogic {
	return &CheckIsFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsFavoriteLogic) CheckIsFavorite(in *userOptPb.CheckIsFavoriteReq) (*userOptPb.CheckIsFavoriteResp, error) {
	result, err := l.svcCtx.UserFavoriteModel.CheckIsFavorite(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		logger.Error("CheckIsFavorite查询错误", err)
		return nil, err
	}
	return &userOptPb.CheckIsFavoriteResp{IsFavorite: result}, nil
}
