package user

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/user/rpc/userservice"

	"douyin/pkg/user/api/internal/svc"
	"douyin/pkg/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	info, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("get user info failed: %v", err.Error())
		return &types.UserInfoRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get user info failed",
			},
		}, nil
	}

	return &types.UserInfoRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		User: &types.User{
			UserId:        info.User.UserId,
			UserName:      info.User.UserName,
			FollowCount:   info.User.FollowCount,
			FollowerCount: info.User.FollowerCount,
			IsFollow:      false,
		},
	}, nil

	return
}
