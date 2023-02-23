package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.FollowerListReq) (resp *types.FollowerListRes, err error) {
	result, err := l.svcCtx.FollowRPC.GetFollowerList(l.ctx, &followservice.GetFollowerListReq{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("UserFollowerList->followerRpc  err : %v , val : %s , message:%+v", err)
		return &types.FollowerListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to followRpc err",
			},
		}, nil
	}
	userList := make([]*types.User, len(result.UserList))
	for i, v := range result.UserList {
		userList[i] = &types.User{}
		userList[i].UserId = v.Id
		userList[i].UserName = v.Name
		userList[i].FollowCount = v.FollowCount
		userList[i].FollowerCount = v.FollowerCount
		userList[i].IsFollow = v.IsFollow
		userList[i].Avatar = v.Avatar
		userList[i].BackgroundImage = v.BackgroundImage
		userList[i].Signature = v.Signature
		userList[i].TotalFavorited = v.TotalFavorited
		userList[i].WorkCount = v.WorkCount
		userList[i].FavoriteCount = v.FavoriteCount
	}
	return &types.FollowerListRes{
		Status: types.Status{
			Code: xerr.OK,
			Msg:  "Get favorite list success",
		},
		UserFollowerlist: userList,
	}, nil
}
