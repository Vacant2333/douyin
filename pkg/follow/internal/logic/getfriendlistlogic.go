package logic

import (
	"context"
	"douyin/pkg/logger"
	"sync"

	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *follow.GetFriendListReq) (*follow.GetFriendListResp, error) {
	followerResp, err := GetFollowerList(l.ctx, l.svcCtx, &follow.GetFollowerListReq{
		Token:  in.Token,
		UserId: in.UserId,
	})

	if err != nil {
		logger.Error("GetFriendList call GetFollowerList failed %s ", err.Error())
		return &follow.GetFriendListResp{
			StatusCode: -1,
			StatusMsg:  "GetFriendList call GetFollowerList failed",
		}, err
	}
	var wg sync.WaitGroup
	userList := make([]*follow.FriendUser, len(followerResp.UserList))
	for index, queryUser := range followerResp.UserList {
		i := index
		user := queryUser
		wg.Add(1)
		go func() {
			userList[i] = &follow.FriendUser{}
			userList[i].Id = user.Id
			userList[i].Name = user.Name
			userList[i].FollowCount = user.FollowCount
			userList[i].FollowerCount = user.FollowerCount
			userList[i].IsFollow = user.IsFollow
			userList[i].Avatar = user.Avatar
			userList[i].BackgroundImage = user.BackgroundImage
			userList[i].Signature = user.Signature
			userList[i].TotalFavorited = user.TotalFavorited
			userList[i].WorkCount = user.WorkCount
			userList[i].FavoriteCount = user.FavoriteCount

			msg, err := l.svcCtx.FollowModel.FindMsg(l.ctx, in.UserId, user.Id)
			if err != nil {
				logger.Errorf("FindMsg failed %s", err.Error())
				return
			}
			defaultMsg := "暂未发消息"
			if msg == nil {
				userList[i].Message = &defaultMsg
			} else {
				userList[i].Message = &msg.Msg.String
				if msg.Sender == 0 {
					userList[i].MsgType = 1
				}
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	return &follow.GetFriendListResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	}, nil
}
