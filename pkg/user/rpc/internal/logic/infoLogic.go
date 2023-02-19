package logic

import (
	"context"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/pkg/user/rpc/internal/svc"
	"douyin/pkg/user/rpc/userInfoPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *userInfoPb.UserInfoReq) (*userInfoPb.UserInfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil {
		logx.Errorf("get user info failed: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "req: %+v", in)
	}
	// 统计关注数和粉丝数
	followNum, err := l.svcCtx.FollowModel.CountByFollowRelation(l.ctx, in.UserId, "fun_id")
	followerNum, err := l.svcCtx.FollowModel.CountByFollowRelation(l.ctx, in.UserId, "user_id")

	var user userInfoPb.User
	user.FollowCount = followNum
	user.FollowerCount = followerNum
	user.UserId = userInfo.Id
	user.UserName = userInfo.Username

	return &userInfoPb.UserInfoResp{
		User: &user,
	}, nil
}
