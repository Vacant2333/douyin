package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/userOptPb"
	_ "douyin/pkg/favorite/userOptPb"
	"douyin/pkg/logger"
	"douyin/pkg/user/userservice"
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
		logger.Fatal("FindManyByTime failed", err)
		return nil, err
	}

	parseToken := token.ParseToken{}
	tokenResult, err := parseToken.ParseToken(in.Token)
	hasUserId := true
	if err != nil || tokenResult.UserId == 0 {
		hasUserId = false // Token解析错误
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

			info, err := l.svcCtx.UserPRC.Info(l.ctx, &userservice.UserInfoReq{
				UserId: query.AuthorId,
			})
			if err != nil {
				logger.Fatal("Video获取Userinfo出错", err)
				return
			}

			videos[i].Author = &video.User{
				Id:              query.AuthorId,
				Name:            info.User.UserName,
				FollowCount:     info.User.FollowCount,
				FollowerCount:   info.User.FollowerCount,
				Avatar:          info.User.Avatar,
				BackgroundImage: info.User.BackgroundImage,
				Signature:       info.User.Signature,
				TotalFavorited:  info.User.TotalFavorited,
				WorkCount:       info.User.WorkCount,
				FavoriteCount:   info.User.FavoriteCount,
			}
			if hasUserId {
				// todo: 调用Follow RPC,查看是否关注了这个人,填入IsFollow
				videos[i].Author.IsFollow = true // 对应函数

				// 调用RPC查看视频是否点赞
				favoriteResp, err := l.svcCtx.FavoritePRC.CheckIsFavorite(l.ctx, &userOptPb.CheckIsFavoriteReq{
					UserId:  tokenResult.UserId,
					VideoId: query.Id,
				})
				if err != nil {
					logger.Fatal("查询视频是否点赞失败", err)
					return
				}
				videos[i].IsFavorite = favoriteResp.IsFavorite
			}
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
