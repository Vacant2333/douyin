package publish

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"douyin/pkg/video/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublishVideoListLogic) GetPublishVideoList(req *types.GetPubVideoListReq) (resp *types.GetPubVideoListRes, err error) {
	videosResp, err := l.svcCtx.VideoRPC.GetAllVideoByUserId(l.ctx, &video.GetAllVideoByUserIdReq{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("get videos by userID failed: %v", err.Error())
		return &types.GetPubVideoListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get videos failed",
			},
		}, err
	}
	videos := make([]*types.PubVideo, len(videosResp.VideoList))
	for i, v := range videosResp.VideoList {
		videos[i].Id = v.Id
		videos[i].User = types.User{
			UserId:          v.Author.Id,
			UserName:        v.Author.Name,
			FollowCount:     v.Author.FollowCount,
			FollowerCount:   v.Author.FollowerCount,
			IsFollow:        v.Author.IsFollow,
			Avatar:          v.Author.Avatar,
			BackgroundImage: v.Author.BackgroundImage,
			Signature:       v.Author.Signature,
			TotalFavorited:  v.Author.TotalFavorited,
			WorkCount:       v.Author.WorkCount,
			FavoriteCount:   v.Author.FavoriteCount,
		}
		videos[i].PlayURL = v.PlayUrl
		videos[i].CoverURL = v.CoverUrl
		videos[i].FavoriteCount = int(v.FavoriteCount)
		videos[i].CommentCount = int(v.CommentCount)
		videos[i].IsFavorite = v.IsFavorite
		videos[i].Title = v.Title
	}
	return &types.GetPubVideoListRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		VideoPubList: videos,
	}, nil
}
