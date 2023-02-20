package logic

import (
	"context"
	"douyin/common/model/videoModel"
	"sync"

	"douyin/pkg/video/rpc/internal/svc"
	"douyin/pkg/video/rpc/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllVideoByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllVideoByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllVideoByUserIdLogic {
	return &GetAllVideoByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllVideoByUserIdLogic) GetAllVideoByUserId(in *video.GetAllVideoByUserIdReq) (*video.GetAllVideoByUserIdResp, error) {
	queryVideos, err := l.svcCtx.VideoModel.FindAllByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	videos := make([]*video.Video, len(queryVideos))
	var wg sync.WaitGroup
	for i, queryVideo := range queryVideos {
		query := queryVideo
		wg.Add(1)
		go func(i int, queryVideos []*videoModel.Video) {
			videos[i] = &video.Video{
				Id:            query.Id,
				PlayUrl:       query.PlayUrl,
				CoverUrl:      query.CoverUrl,
				FavoriteCount: query.LikeCount,
				CommentCount:  query.CommentCount,
				Title:         query.Title,
			}
			// todo: 调用RPC获取作者信息，以及user是否关注了该人
			videos[i].Author = &video.User{
				Id:              in.UserId,
				Name:            "Author",
				FollowCount:     12,
				FollowerCount:   34,
				IsFollow:        false,
				Avatar:          "",
				BackgroundImage: "",
				Signature:       "this is a sign",
				TotalFavorited:  16,
				WorkCount:       0,
				FavoriteCount:   15,
			}
			// todo: 调用RPC获取是否点赞
			videos[i].IsFavorite = false
			defer wg.Done()
		}(i, queryVideos)
	}
	wg.Wait()
	return &video.GetAllVideoByUserIdResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
	}, nil
}
