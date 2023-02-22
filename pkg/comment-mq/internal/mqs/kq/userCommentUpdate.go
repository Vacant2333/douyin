package kq

import (
	"context"
	"douyin/common/messageTypes"
	"douyin/pkg/comment-mq/internal/svc"
	"douyin/pkg/comment/usercomment"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
Listening to the payment flow status change notification message queue
*/
type UserCommentOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCommentUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserCommentOpt {
	return &UserCommentOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCommentOpt) Consume(_, val string) error {
	var message messageTypes.UserCommentOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserCommentOpt) execService(message messageTypes.UserCommentOptMessage) error {
	fmt.Printf("消费者开始消费------------------------------\n")
	if message.ActionType != messageTypes.ActionADD && message.ActionType != messageTypes.ActionCancel {
		return errors.New("UserCommentOptMessage->execService getActionType err")
	}

	// 调用rpc 更新user_comment表
	_, err := l.svcCtx.UserCommentRpc.UpdateCommentStatus(l.ctx, &usercomment.UpdateCommentStatusReq{
		VideoId:    message.VideoId,
		UserId:     message.UserId,
		Content:    message.CommentText,
		CommentId:  message.CommentId,
		ActionType: message.ActionType,
	})

	logx.Error("UserCommentOptMessage->execService xxxxxxxxxxx")

	if err != nil {
		logx.Errorf("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, message)
		return err
	}
	fmt.Printf("消费者消费成功------------------------------\n")
	return nil
}
