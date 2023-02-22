package kq

import (
	"context"
	"douyin/common/messageTypes"
	"douyin/pkg/favorite/useroptservice"
	"douyin/pkg/mq/internal/svc"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
Listening to the payment flow status change notification message queue
*/
type UserFavoriteOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFavoriteUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFavoriteOpt {
	return &UserFavoriteOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavoriteOpt) Consume(_, val string) error {
	var message messageTypes.UserFavoriteOptMessage

	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		logx.Errorf("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}
	return nil
}

// 处理逻辑
func (l *UserFavoriteOpt) execService(message messageTypes.UserFavoriteOptMessage) error {
	fmt.Printf("消费者开始消费------------------------------\n")
	if message.ActionType != messageTypes.ActionADD && message.ActionType != messageTypes.ActionCancel {
		return errors.New("UserCommentOptMessage->execService getActionType err")
	}

	_, err := l.svcCtx.UserFavoriteRpc.UpdateFavoriteStatus(l.ctx, &useroptservice.UpdateFavoriteStatusReq{
		VideoId:    message.VideoId,
		UserId:     message.UserId,
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
