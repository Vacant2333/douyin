package userOpt

import (
	"context"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/logger"
	"sync"

	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.FriendListReq) (resp *types.FriendListRes, err error) {
	friendRes, err := l.svcCtx.FollowRPC.GetFriendList(l.ctx, &followservice.GetFriendListReq{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		logger.Error("GetFriendList RPC failed %s", err.Error())
		return &types.FriendListRes{
			Status: types.Status{
				Code: -1,
				Msg:  "get friend list failed",
			},
		}, err
	}
	userList := make([]*types.FriendUser, len(friendRes.UserList))
	var wg sync.WaitGroup
	for index, queryUser := range friendRes.UserList {
		i := index
		user := queryUser
		wg.Add(1)
		go func() {
			userList[i] = &types.FriendUser{}
			userList[i].Message = *user.Message
			userList[i].MsgType = user.MsgType
			userList[i].UserId = user.Id
			userList[i].UserName = user.Name
			userList[i].FollowCount = user.FollowCount
			userList[i].FollowerCount = user.FollowerCount
			userList[i].IsFollow = user.IsFollow
			userList[i].Avatar = user.Avatar
			userList[i].BackgroundImage = user.BackgroundImage
			userList[i].Signature = user.Signature
			userList[i].TotalFavorited = user.TotalFavorited
			userList[i].WorkCount = user.WorkCount
			userList[i].FavoriteCount = user.FavoriteCount
			defer wg.Done()
		}()
	}
	wg.Wait()
	return &types.FriendListRes{
		Status: types.Status{
			Code: 0,
			Msg:  "success",
		},
		UserFriendlist: userList,
	}, nil
}
