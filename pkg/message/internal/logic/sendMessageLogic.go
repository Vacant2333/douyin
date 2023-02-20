package logic

import (
	"context"
	"douyin/common/model/messageModel"
	"douyin/pkg/message/internal/svc"
	"douyin/pkg/message/userMessagePb"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------SendMessage-----------------------
func (l *SendMessageLogic) SendMessage(in *userMessagePb.MessageReq) (*userMessagePb.MessageRes, error) {
	fmt.Printf(":::::::::::::::::::::::::::::::::::::::::::::::")
	var message *messageModel.Message
	message.ToUserId = in.ToUserId
	message.Content = in.Content
	message.CreateTime = time.Now()

	_, err := l.svcCtx.MessageModel.Insert(l.ctx, message)

	if err != nil {
		logx.Errorf("SendMessage------->SELECT err : %s", err.Error())
		return &userMessagePb.MessageRes{
			Code: 1,
			Msg:  "failure",
		}, err
	}

	return &userMessagePb.MessageRes{
		Code: 0,
		Msg:  "Success",
	}, nil
}
