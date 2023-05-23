package userOpt

import (
	"context"
	"douyin/pkg/favorite/common/xerr"
	"douyin/pkg/favorite/rpc/useroptservice"

	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"

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
	_, err = l.svcCtx.UserFavoriteRpc.UpdateFavoriteStatus(l.ctx, &useroptservice.UpdateFavoriteStatusReq{
		VideoId:    req.VideoId,
		UserId:     1,
		ActionType: req.ActionType,
	})
	if err != nil {
		logx.Errorf("UserFavoriteOpt->UserFavoriteRpc  err : %v , val : %s , message:%+v", err)
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to UserFavoriteRpc err",
			},
		}, err
	}

	if req.ActionType == 1 {
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user favorite success",
			},
		}, nil
	}
	if req.ActionType == 2 {
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user cancel favorite success",
			},
		}, nil
	}
	return nil, nil
}
