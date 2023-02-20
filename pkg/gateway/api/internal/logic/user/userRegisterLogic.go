package user

import (
	"context"
	"douyin/common/checkpwd"
	"douyin/common/xerr"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"
	"douyin/pkg/user/rpc/userservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterRes, err error) {
	err = checkpwd.CheckPassword(req.Password)
	if err != nil {
		return &types.UserRegisterRes{
			Status: types.Status{
				Code: xerr.REUQEST_PARAM_ERROR,
				Msg:  "Password strength must be the word write + number + symbol, more than 9 digits",
			},
		}, nil
	}

	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userservice.RegisterReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("register failed: %s", err.Error())
		return &types.UserRegisterRes{
			Status: types.Status{
				Code: xerr.SECRET_ERROR,
				Msg:  "register failed" + err.Error(),
			},
		}, nil
	}

	return &types.UserRegisterRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		IdWithTokenRes: types.IdWithTokenRes{
			UserId: res.UserId,
			Token:  res.Token,
		},
	}, nil
}
