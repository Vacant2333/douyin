package logic

import (
	"context"
	"fmt"

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
	fmt.Printf("GetVideoFavorite-------------->")

	allFavoriteInfoData, err := l.svcCtx.UserFavoriteModel.FindAll(l.ctx, in.UserId)
	//fmt.Printf("favoriteList: %v", favoriteList)
	if err != nil {
		logx.Errorf("GetFavoriteList------->SELECT err : %s", err.Error())
		return &userOptPb.GetUserFavoriteResp{}, err
	}
	var favoriteList []*userOptPb.Favorite
	for _, v := range allFavoriteInfoData {
		var favorite userOptPb.Favorite
		videoInfoData, err := l.svcCtx.UserVideoModel.FindOne(l.ctx, v.VideoId)
		if err != nil {
			logx.Errorf("GetFavoriteList------->SELECT err : %s", err.Error())
			return &userOptPb.GetUserFavoriteResp{}, err
		}
		favorite.PlayUrl = videoInfoData.PlayUrl
		favorite.Title = videoInfoData.Title
		favorite.VideoId = videoInfoData.Id
		favorite.CoverUrl = videoInfoData.CoverUrl
		favorite.AuthorId = videoInfoData.AuthorId

		favoriteList = append(favoriteList, &favorite)
	}
	return &userOptPb.GetUserFavoriteResp{
		FavoriteList: favoriteList,
	}, nil
}
