package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/user/internal/svc"
	"douyin/pkg/user/userInfoPb"

	"github.com/pkg/errors"

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
	check(err)
	followerNum, err := l.svcCtx.FollowModel.CountByFollowRelation(l.ctx, in.UserId, "user_id")
	check(err)
	// 统计视频数
	workCount, err := l.svcCtx.FavoriteModel.FindAllByUserId(l.ctx, in.UserId)
	check(err)

	// 统计关注情况
	favoriteCount, err := l.svcCtx.FavoriteModel.FindAllByUserId(l.ctx, in.UserId)
	check(err)

	var user userInfoPb.User
	user.FollowCount = followNum
	user.FollowerCount = followerNum
	user.UserId = userInfo.Id
	user.UserName = userInfo.Username
	user.Avatar = userInfo.Avatar
	user.BackgroundImage = userInfo.BackgroundImage
	user.Signature = userInfo.Signature
	user.FavoriteCount = favoriteCount
	user.WorkCount = workCount
	user.TotalFavorited = "测试"

	return &userInfoPb.UserInfoResp{
		User: &user,
	}, nil
}

func check(e error) {
	if e != nil {
		logx.Errorf(e.Error())
	}
}
