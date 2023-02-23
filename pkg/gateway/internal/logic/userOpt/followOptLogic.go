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

type FollowOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowOptLogic {
	return &FollowOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowOptLogic) FollowOpt(req *types.FollowOptReq) (resp *types.FollowOptRes, err error) {
	var followTemp messageTypes.UserFollowOptMessage
	_ = copier.Copy(&followTemp, req)

	followTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
	followTemp.ToUserId = req.FollowId
	followTemp.ActionType = req.ActionType
	// 序列化
	msg, err := json.Marshal(followTemp)
	if err != nil {
		logx.Errorf("FollowOpt json.Marshal err : %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), " json.Marshal err")
	}
	// 向消息队列发送消息
	err = l.svcCtx.FollowOptMsgProducer.Push(string(msg))
	if err != nil {
		logx.Errorf("FollowOpt msgProducer.Push err : %s", err.Error())
		return &types.FollowOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to FollowOptMsgConsumer err",
			},
		}, nil
	}
	if req.ActionType == 1 {
		return &types.FollowOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user follow success",
			},
		}, nil
	}
	if req.ActionType == 2 {
		return &types.FollowOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user cancel follow success",
			},
		}, nil
	}
	return nil, nil
}
