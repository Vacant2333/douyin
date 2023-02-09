package logic

import (
	"context"

	"douyin/pkg/userinfo-demo/rpc/internal/svc"
	"douyin/pkg/userinfo-demo/rpc/types/userinfo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *userinfo.UserinfoReq) (*userinfo.UserinfoRes, error) {
	// todo: add your logic here and delete this line

	return &userinfo.UserinfoRes{}, nil
}
