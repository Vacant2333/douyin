package message

import (
	"context"
	"douyin/pkg/message/userMessagePb"

	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessagelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessagelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessagelistLogic {
	return &MessagelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessagelistLogic) Messagelist(req *types.MessageListReq) (resp *types.MessageListRes, err error) {
	GetMessagelistRPC, err := l.svcCtx.MessageRpc.GetMessageList(l.ctx, &userMessagePb.MessageListReq{UserId: req.UserId, Token: req.Token})
	if err != nil {
		logx.Errorf("GetMessagelist->GetMessagelistRpc  err : %v , val : %s , message:%+v", err)
		return &types.MessageListRes{
			Status:      types.Status{Code: 0, Msg: GetMessagelistRPC.Msg},
			MessageList: nil,
		}, err
	}

	var MessageList []*types.Message

	for _, v := range GetMessagelistRPC.MessageList {
		var message *types.Message
		message.Id = v.Id
		message.Content = v.Content
		message.CreateTime = v.CreateTime
		message.FromUserId = v.FromUserId
		message.ToUserId = v.ToUserId

		MessageList = append(MessageList, message)
	}
	return &types.MessageListRes{
		Status:      types.Status{Code: 0, Msg: GetMessagelistRPC.Msg},
		MessageList: MessageList,
	}, nil
}
