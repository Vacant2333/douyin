package logic

import (
	"context"
	"douyin/common/model/videoModel"
	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoLogic) GetVideo(in *video.GetVideoReq) (*video.GetVideoResp, error) {
	// todo: 这个10应该改为配置文件中读取，待改进
	var selectNum int64 = 10
	queryVideos, err := l.svcCtx.VideoModel.FindManyByTime(l.ctx, in.LatestTime, selectNum)
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
				Id:              "123",
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
	nextTime := queryVideos[len(videos)-1].Time
	return &video.GetVideoResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
		NextTime:   nextTime,
	}, nil
}
