package user

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"douyin/pkg/user/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginRes, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userservice.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("login failed: %v", err.Error())
		return &types.UserLoginRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "login failed",
			},
		}, nil
	}

	return &types.UserLoginRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		IdWithTokenRes: types.IdWithTokenRes{
			UserId: res.UserId,
			Token:  res.Token,
		},
	}, nil
}
