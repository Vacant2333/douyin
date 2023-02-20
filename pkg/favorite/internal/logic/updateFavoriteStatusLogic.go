package logic

import (
	"context"
	"douyin/common/globalkey"
	"douyin/common/messageTypes"
	"douyin/common/model/favoriteModel"
	"douyin/common/xerr"
	"douyin/pkg/favorite/internal/svc"
	"douyin/pkg/favorite/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFavoriteStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFavoriteStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFavoriteStatusLogic {
	return &UpdateFavoriteStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFavoriteStatusLogic) UpdateFavoriteStatus(in *userOptPb.UpdateFavoriteStatusReq) (*userOptPb.UpdateFavoriteStatusResp, error) {
	favoriteModel := new(favoriteModel.Favorite)
	switch in.ActionType {
	//新增点赞
	case messageTypes.ActionADD:

		favoriteModel.UserId = in.UserId
		favoriteModel.VideoId = in.VideoId

		_, err := l.svcCtx.UserFavoriteModel.Insert(l.ctx, favoriteModel)

		if err != nil {
			logx.Errorf("UpdateFavoriteStatus------->Insert err : %s", err.Error())
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}

		return &userOptPb.UpdateFavoriteStatusResp{}, nil

	//取消点赞
	case messageTypes.ActionCancel:
		favoriteModel.Removed = globalkey.DelStateYes
		favoriteModel.VideoId = in.VideoId
		favoriteModel.UserId = in.UserId
		err := l.svcCtx.UserFavoriteModel.Update(l.ctx, favoriteModel)
		if err != nil {
			logx.Errorf("UpdateFavoriteStatus------->update err : %s", err.Error())
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}
		return &userOptPb.UpdateFavoriteStatusResp{}, nil

	default:
		return &userOptPb.UpdateFavoriteStatusResp{}, xerr.NewErrMsg("actionType must be 1 or 2")
	}
}
