package userOpt

import (
	"context"
	"douyin/pkg/comment/common/xerr"
	"douyin/pkg/comment/rpc/usercomment"
	"time"

	"douyin/pkg/comment/api/internal/svc"
	"douyin/pkg/comment/api/internal/types"

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
	//获取参数

	//校验参数

	//敏感词过滤

	// 调用rpc 更新user_comment表
	_, err = l.svcCtx.UserCommentRpc.UpdateCommentStatus(l.ctx, &usercomment.UpdateCommentStatusReq{
		VideoId:    req.VideoId,
		UserId:     1, /*l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)*/
		CommentId:  req.CommentId,
		ActionType: req.ActionType,
	})
	if err != nil {
		logx.Errorf("UserCommentOpt->commentRpc  err : %v , val : %s , message:%+v", err)
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer err",
			},
		}, nil
	}

	if req.ActionType == 1 {
		// 调用user-info rpc拉取发布消息的用户信息

		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			Comment: &types.Comment{
				CommentId:  req.CommentId,
				Content:    req.CommentText,
				CreateTime: time.Now().Format("01-01"),
			},
		}, nil
	}

	if req.ActionType == 2 {
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}

	return nil, nil
}
