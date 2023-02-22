package kq

import (
	"context"
	"douyin/common/messageTypes"
	"douyin/pkg/follow/followservice"
	"douyin/pkg/mq/internal/svc"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
Listening to the payment flow status change notification message queue
*/
type UserFollowOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowOpt {
	return &UserFollowOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowOpt) Consume(_, val string) error {
	var message messageTypes.UserFollowOptMessage

	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserFollowOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserFollowOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		logx.Errorf("UserFollowOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}
	return nil
}

// 处理逻辑
func (l *UserFollowOpt) execService(message messageTypes.UserFollowOptMessage) error {
	if message.ActionType != messageTypes.ActionADD && message.ActionType != messageTypes.ActionCancel {
		return errors.New("UserCommentOptMessage->execService getActionType err")
	}

	_, err := l.svcCtx.UserFollowRPC.Follow(l.ctx, &followservice.FollowReq{
		ToUserId:   message.ToUserId,
		ActionType: message.ActionType,
		UserId:     message.UserId,
	})

	if err != nil {
		logx.Errorf("UserFollowOptMessage->execService  err : %v , val : %s , message:%+v", err, message)
		return err
	}
	fmt.Printf("消费者消费成功------------------------------\n")
	return nil
}
