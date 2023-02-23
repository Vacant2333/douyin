package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/model/messageModel"
	"douyin/pkg/logger"
	"douyin/pkg/message/internal/svc"
	"douyin/pkg/message/userMessagePb"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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

func (l *SendMessageLogic) SendMessage(in *userMessagePb.MessageReq) (*userMessagePb.MessageRes, error) {
	message := &messageModel.Message{}
	var parseToken token.ParseToken
	claims, err := parseToken.ParseToken(in.Token)
	if err != nil {
		logger.InfoF("Token解析错误 %s ", err.Error())
		return &userMessagePb.MessageRes{
			Code: -1,
			Msg:  "Token解析错误",
		}, err
	}

	message.ToUserId = in.ToUserId
	message.Content = in.Content
	message.FromUserId = claims.UserId
	message.CreateTime = time.Now().Unix()

	_, err = l.svcCtx.MessageModel.Insert(l.ctx, message)

	if err != nil {
		logx.Errorf("SendMessage------->SELECT err : %s", err.Error())
		return &userMessagePb.MessageRes{
			Code: -1,
			Msg:  "failure",
		}, err
	}

	err = l.svcCtx.FollowModel.UpdateMsg(l.ctx, claims.UserId, in.ToUserId, message.Content)
	if err != nil {
		logx.Errorf("Update First Msg failed : %s", err.Error())
		return &userMessagePb.MessageRes{
			Code: -1,
			Msg:  "failure",
		}, err
	}

	return &userMessagePb.MessageRes{
		Code: 0,
		Msg:  "Success",
	}, nil
}
