package logic

import (
	"context"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"
	"douyin/pkg/user/userservice"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *follow.GetFollowListReq) (*follow.GetFollowListResp, error) {
	queryFollow, err := l.svcCtx.FollowModel.FindAllByFunId(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("get follow list failed: %v", err.Error())
		return &follow.GetFollowListResp{
			StatusCode: -1,
			StatusMsg:  "get follow list failed",
		}, err
	}
	follows := make([]*follow.User, len(queryFollow))
	var wg sync.WaitGroup
	for index, v := range queryFollow {
		wg.Add(1)
		query := v
		i := index
		go func() {
			queryFollows, err := l.svcCtx.UserPRC.Info(l.ctx, &userservice.UserInfoReq{
				UserId: query.UserId.Int64,
			})
			if err != nil {
				logx.Errorf("in follow model get user info failed: %v", err.Error())
				return
			}
			follows[i].Id = queryFollows.User.UserId
			follows[i].Name = queryFollows.User.UserName
			follows[i].FollowCount = queryFollows.User.FollowCount
			follows[i].FollowerCount = queryFollows.User.FollowerCount
			follows[i].IsFollow = true
			follows[i].Avatar = queryFollows.User.Avatar
			follows[i].BackgroundImage = queryFollows.User.BackgroundImage
			follows[i].Signature = queryFollows.User.Signature
			follows[i].TotalFavorited = queryFollows.User.TotalFavorited
			follows[i].WorkCount = queryFollows.User.WorkCount
			follows[i].FavoriteCount = queryFollows.User.FavoriteCount
			defer wg.Done()
		}()
	}
	wg.Wait()
	return &follow.GetFollowListResp{
		StatusCode: 0,
		UserList:   follows,
	}, nil
}
