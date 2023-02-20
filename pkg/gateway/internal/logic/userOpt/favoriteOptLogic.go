package userOpt

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/messageTypes"
	"douyin/common/xerr"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteOptLogic {
	return &FavoriteOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteOptLogic) FavoriteOpt(req *types.FavoriteOptReq) (resp *types.FavoriteOptRes, err error) {
	var msgTemp messageTypes.UserFavoriteOptMessage
	_ = copier.Copy(&msgTemp, req)

	msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)

	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		logx.Errorf("FavoriteOpt json.Marshal err : %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), " json.Marshal err")
	}

	// 向消息队列发送消息
	err = l.svcCtx.FavoriteOptMsgProducer.Push(string(msg))
	if err != nil {
		logx.Errorf("FavoriteOpt msgProducer.Push err : %s", err.Error())
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to FavoriteOptMsgConsumer err",
			},
		}, nil
	}
	if req.ActionType == 1 {
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user favorite success",
			},
		}, nil
	}
	if req.ActionType == 2 {
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user cancel favorite success",
			},
		}, nil
	}
	return nil, nil
}
