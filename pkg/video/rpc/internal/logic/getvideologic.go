package logic

import (
	"context"
	"douyin/common/model/videoModel"
	"douyin/pkg/video/rpc/internal/svc"
	"douyin/pkg/video/rpc/types/video"
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
	videos := make([]*video.Video, selectNum)
	var wg sync.WaitGroup
	for i, queryVideo := range queryVideos {
		query := queryVideo
		wg.Add(1)
		go func(i int, queryVideos []*videoModel.Video) {
			videos[i].Id = query.Id
			videos[i].PlayUrl = query.PlayUrl
			videos[i].CoverUrl = query.CoverUrl
			videos[i].FavoriteCount = query.LikeCount
			videos[i].CommentCount = query.CommentCount
			videos[i].Title = query.Title
			// todo: 调用RPC获取作者信息
			videos[i].Author = nil
			// todo: 调用RPC获取是否点赞
			videos[i].IsFavorite = true
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
