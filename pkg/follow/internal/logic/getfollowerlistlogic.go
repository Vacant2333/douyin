package logic

import (
	"context"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"
	"douyin/pkg/user/userservice"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *follow.GetFollowerListReq) (*follow.GetFollowerListResp, error) {
	return GetFollowerList(l.ctx, l.svcCtx, in)
}

func GetFollowerList(ctx context.Context, svcCtx *svc.ServiceContext, in *follow.GetFollowerListReq) (*follow.GetFollowerListResp, error) {
	queryFollow, err := svcCtx.FollowModel.FindAllByUserId(ctx, in.UserId)
	if err != nil {
		logx.Errorf("get follower list failed: %v", err.Error())
		return &follow.GetFollowerListResp{
			StatusCode: -1,
			StatusMsg:  "get follower list failed",
		}, err
	}
	followers := make([]*follow.User, len(queryFollow))
	var wg sync.WaitGroup
	for index, v := range queryFollow {
		wg.Add(1)
		query := v
		i := index
		go func() {
			queryFollowers, err := svcCtx.UserPRC.Info(ctx, &userservice.UserInfoReq{
				UserId: query.FunId,
			})
			if err != nil {
				logx.Errorf("in follow model get user info failed: %v", err.Error())
				return
			}
			followers[i] = &follow.User{}
			followers[i].Id = queryFollowers.User.UserId
			followers[i].Name = queryFollowers.User.UserName
			followers[i].FollowCount = queryFollowers.User.FollowCount
			followers[i].FollowerCount = queryFollowers.User.FollowerCount
			followers[i].Avatar = queryFollowers.User.Avatar
			followers[i].BackgroundImage = queryFollowers.User.BackgroundImage
			followers[i].Signature = queryFollowers.User.Signature
			followers[i].TotalFavorited = queryFollowers.User.TotalFavorited
			followers[i].WorkCount = queryFollowers.User.WorkCount
			followers[i].FavoriteCount = queryFollowers.User.FavoriteCount

			isFollow, err := svcCtx.FollowModel.CheckIsFollow(ctx, queryFollowers.User.UserId, in.UserId)
			if err != nil {
				logx.Errorf("in follow model query is follow failed: %v", err.Error())
				return
			}
			followers[i].IsFollow = isFollow
			defer wg.Done()
		}()
	}
	wg.Wait()
	return &follow.GetFollowerListResp{
		StatusCode: 0,
		UserList:   followers,
	}, nil
}
