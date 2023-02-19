package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/favorite/rpc/userOptPb"
	"douyin/pkg/user/rpc/userservice"
	"fmt"

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
	fmt.Printf("len::::::::::::::::::::::::::::::::::::::::%v\n", len(result.FavoriteList))
	for _, v := range result.FavoriteList {
		var Favorite types.PubVideo
		userinfo, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.UserInfoReq{
			UserId: v.AuthorId,
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
		//评论列表
		allCommentInfoData, err := l.svcCtx.UserCommentRpc.GetVideoComment(l.ctx, &usercomment.GetVideoCommentReq{
			VideoId: v.VideoId,
		})
		if err != nil {
			logx.Errorf("UserCommentList->commentRpc  err : %v , val : %s , message:%+v", err)
			return &types.FavoriteListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "send message to commentRpc err",
				},
			}, err
		}

		var author = types.User{
			UserId:        userinfo.User.UserId,
			UserName:      userinfo.User.UserName,
			FollowCount:   userinfo.User.FollowCount,
			FollowerCount: userinfo.User.FollowerCount,
			IsFollow:      false,
		}
		Favorite.User = author
		Favorite.IsFavorite = true
		Favorite.FavoriteCount = len(result.FavoriteList)
		Favorite.Title = v.Title
		Favorite.CommentCount = len(allCommentInfoData.CommentList)
		Favorite.CoverURL = v.CoverUrl
		Favorite.PlayURL = v.PlayUrl
		FavoriteList = append(FavoriteList, &Favorite)
	}

	return &types.FavoriteListRes{
		Status: types.Status{
			Code: xerr.OK,
			Msg:  "Get favorite list success",
		},
		FavoriteList: FavoriteList,
	}, nil
}
