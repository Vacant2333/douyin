package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/model/videoModel"
	"douyin/pkg/favorite/userOptPb"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/logger"
	"douyin/pkg/user/userservice"
	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
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

	authorId := in.UserId
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

				checkIsFollowResp, err := l.svcCtx.FollowRPC.CheckIsFollow(l.ctx, &followservice.CheckIsFollowReq{
					UserId: authorId,
					FunId:  in.UserId,
				})
				if err != nil {
					logger.Fatalf("CheckIsFollow RPC failed %s", err.Error())
					return
				}
				videos[i].Author.IsFollow = checkIsFollowResp.IsFollow

				// 调用RPC查看视频是否点赞
				favoriteResp, err := l.svcCtx.FavoriteRPC.CheckIsFavorite(l.ctx, &userOptPb.CheckIsFavoriteReq{
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
