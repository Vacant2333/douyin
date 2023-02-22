package message

import (
	"context"
	"douyin/common/help/sensitiveWords"
	"douyin/pkg/message/userMessagePb"

	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageLogic) Message(req *types.MessageReq) (resp *types.MessageRes, err error) {
	req.Content = sensitiveWords.SensitiveWordsFliter(sensitiveWords.SensitiveWords, req.Content, '?')
	SendMessageRPC, err := l.svcCtx.MessageRpc.SendMessage(l.ctx, &userMessagePb.MessageReq{ToUserId: req.ToUserId, Content: req.Content, Token: req.Token, ActionType: req.ActionType})

	if err != nil {
		logx.Errorf("SendMessage->SendMessageRpc  err : %v , val : %s , message:%+v", err)
		return &types.MessageRes{
			Code: 1,
			Msg:  "send message to SendMessageRPC err",
		}, err
	}
	return &types.MessageRes{
		Code: 0,
		Msg:  SendMessageRPC.Msg,
	}, nil
}
