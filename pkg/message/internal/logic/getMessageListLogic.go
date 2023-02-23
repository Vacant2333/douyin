package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/pkg/message/internal/svc"
	"douyin/pkg/message/userMessagePb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------GetMessageList-----------------------
func (l *GetMessageListLogic) GetMessageList(in *userMessagePb.MessageListReq) (*userMessagePb.MessageListRes, error) {
	parseToken := token.ParseToken{}
	userid, _ := parseToken.ParseToken(in.Token)
	allMessageData, err := l.svcCtx.MessageModel.FindMessageListByUserID(l.ctx, userid.UserId, in.UserId, in.PreTime)
	if err != nil {
		logx.Errorf("GetCommentList------->SELECT err : %s", err.Error())
		return &userMessagePb.MessageListRes{
			Code:        1,
			Msg:         "failure",
			MessageList: nil,
		}, err
	}

	var MessageList []*userMessagePb.Message
	for _, v := range allMessageData {
		var message userMessagePb.Message
		message.Id = v.Id
		message.ToUserId = v.ToUserId
		message.Content = v.Content
		message.FromUserId = v.FromUserId
		message.CreateTime = v.CreateTime
		MessageList = append(MessageList, &message)
	}
	return &userMessagePb.MessageListRes{
		Code:        0,
		Msg:         "Success",
		MessageList: MessageList,
	}, nil
}
