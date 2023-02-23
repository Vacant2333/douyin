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
	newFavoriteModel := new(favoriteModel.Favorite)
	switch in.ActionType {
	//新增点赞
	case messageTypes.ActionADD:
		newFavoriteModel.UserId = in.UserId
		newFavoriteModel.VideoId = in.VideoId

		_, err := l.svcCtx.UserFavoriteModel.Insert(l.ctx, newFavoriteModel)

		if err != nil {
			logx.Errorf("UpdateFavoriteStatus------->Insert err : %s", err.Error())
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}

		//_, err = l.svcCtx.VideoRPC.ChangeVideoFavorite(l.ctx, &videoservice.ChangeVideoFavoriteReq{
		//	VideoId:    in.VideoId,
		//	ActionType: in.ActionType,
		//})
		err = l.svcCtx.VideoModel.UpdateCount(l.ctx, in.VideoId, "like_count", in.ActionType)
		if err != nil {
			logx.Errorf("ChangeVideoFavorite failed %s ", err)
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}

		return &userOptPb.UpdateFavoriteStatusResp{}, nil

	//取消点赞
	case messageTypes.ActionCancel:
		newFavoriteModel.Removed = globalkey.DelStateYes
		newFavoriteModel.VideoId = in.VideoId
		newFavoriteModel.UserId = in.UserId
		err := l.svcCtx.UserFavoriteModel.Update(l.ctx, newFavoriteModel)
		if err != nil {
			logx.Errorf("UpdateFavoriteStatus------->update err : %s", err.Error())
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}

		//_, err = l.svcCtx.VideoRPC.ChangeVideoFavorite(l.ctx, &videoservice.ChangeVideoFavoriteReq{
		//	VideoId:    in.VideoId,
		//	ActionType: in.ActionType,
		//})
		err = l.svcCtx.VideoModel.UpdateCount(l.ctx, in.VideoId, "like_count", in.ActionType)

		if err != nil {
			logx.Errorf("ChangeVideoFavorite failed %s ", err)
			return &userOptPb.UpdateFavoriteStatusResp{}, err
		}

		return &userOptPb.UpdateFavoriteStatusResp{}, nil

	default:
		return &userOptPb.UpdateFavoriteStatusResp{}, xerr.NewErrMsg("actionType must be 1 or 2")
	}
}
