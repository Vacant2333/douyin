package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/userOptPb"
	"douyin/pkg/user/userservice"
	"strconv"
	"sync"

	"douyin/pkg/logger"
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
	// 根据token查看是否已登陆
	parseToken := token.ParseToken{}
	tokenResult, err := parseToken.ParseToken(in.Token)
	hasUserId := true
	if err != nil || tokenResult.UserId == 0 {
		hasUserId = false // Token解析错误
	}

	id, err := strconv.Atoi(in.UserId)
	authorId := int64(id)
	if err != nil {
		logger.Fatal("useId转换int类型失败", err)
		return nil, err
	}
	queryVideos, err := l.svcCtx.VideoModel.FindAllByUserId(l.ctx, authorId)
	if err != nil {
		logger.Fatal("FindAllByUserId failed", err)
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
			// 查询作者信息
			info, err := l.svcCtx.UserPRC.Info(l.ctx, &userservice.UserInfoReq{
				UserId: authorId,
			})
			if err != nil {
				logger.Fatal("Video获取Userinfo出错", err)
				return
			}
			// todo: 调用Follow RPC,查看是否关注了这个人,填入IsFollow
			videos[i].Author = &video.User{
				Id:              authorId,
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
				// todo: 调用PRC查看是否已关注
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
	return &video.GetAllVideoByUserIdResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
	}, nil
}
