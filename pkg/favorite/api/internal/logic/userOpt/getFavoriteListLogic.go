package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/favorite/rpc/userOptPb"

	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteListLogic {
	return &GetFavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFavoriteListLogic) GetFavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListRes, err error) {
	//获取参数
	// 调用rpc 更新user_favorite表
	result, err := l.svcCtx.UserFavoriteRpc.GetUserFavorite(l.ctx, &userOptPb.GetUserFavoriteReq{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("UserFavoriteList->favoriteRpc  err : %v , val : %s , message:%+v", err)
		return &types.FavoriteListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to favoriteRpc err",
			},
		}, nil
	}
	var FavoriteList []*types.PubVideo
	for _, v := range result.FavoriteList {
		var Favorite types.PubVideo
		userinfo, err := l.svcCtx.UserInfoRpc.GetUser(l.ctx, &userinfo.UserinfoRequest{
			UserId: string(v.AuthorId),
		})
		if err != nil {
			logx.Errorf("userinfo->userinfoRpc  err : %v , val : %s , message:%+v", err)
			return &types.FavoriteListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "send message to UserInfoRpc err",
				},
			}, err
		}
		var author = types.User{
			UserId:        1,
			UserName:      userinfo.User.Name,
			FollowCount:   userinfo.User.FollowCount,
			FollowerCount: userinfo.User.FollowerCount,
			IsFollow:      false,
		}
		Favorite.User = author
		Favorite.IsFavorite = true
		Favorite.FavoriteCount = 1
		Favorite.Title = v.Title
		Favorite.CommentCount = 1
		Favorite.CoverURL = v.CoverUrl
		Favorite.PlayURL = v.PlayUrl
		FavoriteList = append(FavoriteList, &Favorite)
	}

	return &types.FavoriteListRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		FavoriteList: FavoriteList,
	}, nil
}
