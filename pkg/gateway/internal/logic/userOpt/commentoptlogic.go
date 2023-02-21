package userOpt

import (
	"context"
	"douyin/common/help/sensitiveWords"
	myToken "douyin/common/help/token"
	"douyin/common/messageTypes"
	"douyin/common/xerr"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"douyin/pkg/user/userservice"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentOptLogic {
	return &CommentOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentOptLogic) CommentOpt(req *types.CommentOptReq) (resp *types.CommentOptRes, err error) {
	// 前端传入的是1，2表示评论取消评论，入口这里就将它转换成1，0表示评论取消评论
	msgTemp, status, err := l.getActionType(req)

	if msgTemp.ActionType == messageTypes.ActionErr || err != nil {
		logx.Errorf("CommentOptLogic CommentOpt err: %s", err.Error())
		return status, nil
	}

	if err != nil {
		// FIXME: args error
		logx.Errorf("UserCommentOpt->UserCommentRpc  err : %v " /* val : %s , message:%+v"*/, err)
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to UserCommentRpc err",
			},
		}, nil
	}

	if req.ActionType == 1 {
		// 调用user-info rpc拉取发布消息的用户信息
		var userId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		userInfoResult, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.UserInfoReq{UserId: userId})
		if err != nil {
			// FIXME: args error
			logx.Errorf("UserCommentOpt->userInfoRpc  err : %v " /*val : %s , message:%+v"*/, err)
			return &types.CommentOptRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "send message to userInfoRpc err",
				},
			}, nil
		}

		var user = types.User{
			UserId:        l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64),
			UserName:      userInfoResult.User.UserName,
			FollowCount:   userInfoResult.User.FollowCount,
			FollowerCount: userInfoResult.User.FollowerCount,
			IsFollow:      false,
		}

		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user comment opt（comment） success",
			},
			Comment: &types.Comment{
				CommentId:  msgTemp.CommentId,
				Content:    msgTemp.CommentText,
				User:       user,
				CreateTime: time.Now().Format("01-01"),
			},
		}, nil
	}

	if req.ActionType == 2 {
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
				Msg:  "user comment opt（cancel comment） success",
			},
		}, nil
	}

	return nil, nil
}

func (l *CommentOptLogic) getActionType(req *types.CommentOptReq) (*messageTypes.UserCommentOptMessage, *types.CommentOptRes, error) {
	var msgTemp messageTypes.UserCommentOptMessage
	_ = copier.Copy(&msgTemp, req)
	fmt.Printf("userid:::::::%v", l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64))

	switch req.ActionType { // 方便扩展
	case messageTypes.ActionADD:
		//敏感词过滤
		msgTemp.CommentText = sensitiveWords.SensitiveWordsFliter(sensitiveWords.SensitiveWords, msgTemp.CommentText, '?')
		msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		//l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		msgTemp.CreateDate = time.Now().Format("01-01")
		msgTemp.ActionType = messageTypes.ActionADD

	case messageTypes.ActionCancel:
		msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		//l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		msgTemp.ActionType = messageTypes.ActionCancel
	default:
		msgTemp.ActionType = -99
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer ActionType err",
			},
		}, errors.New("operate error")
	}

	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer json.Marshal err",
			},
		}, errors.Wrapf(err, " json.Marshal err")
	}

	// 向消息队列发送消息
	err = l.svcCtx.CommentOptMsgProducer.Push(string(msg))
	if err != nil {
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer err",
			},
		}, errors.Wrapf(err, " json.Marshal err")
	}
	return &msgTemp, nil, nil
}
