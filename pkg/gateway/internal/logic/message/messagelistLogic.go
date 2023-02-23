package message

import (
	"context"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"douyin/pkg/message/userMessagePb"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessagelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageListLogin(ctx context.Context, svcCtx *svc.ServiceContext) *MessagelistLogic {
	return &MessagelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessagelistLogic) MessageList(req *types.MessageListReq) (resp *types.MessageListRes, err error) {
	GetMessagelistRPC, err := l.svcCtx.MessageRpc.GetMessageList(l.ctx, &userMessagePb.MessageListReq{
		UserId:  req.ToUserId,
		Token:   req.Token,
		PreTime: req.PreMsgTime,
	})
	if err != nil {
		logx.Errorf("GetMessagelist->GetMessagelistRpc  err : %v , val : %s , message:%+v", err)
		return &types.MessageListRes{
			Status:      types.Status{Code: 1, Msg: err.Error()},
			MessageList: nil,
		}, err
	}
	MessageList := make([]*types.Message, len(GetMessagelistRPC.MessageList))

	for i, v := range GetMessagelistRPC.MessageList {
		message := &types.Message{}
		message.Id = v.Id
		message.Content = v.Content
		message.CreateTime = v.CreateTime
		message.FromUserId = v.FromUserId
		message.ToUserId = v.ToUserId
		MessageList[i] = message
	}
	return &types.MessageListRes{
		Status:      types.Status{Code: 0},
		MessageList: MessageList,
	}, nil
}
