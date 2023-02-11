package logic

import (
	"comment/rpc/internal/svc"
	"comment/rpc/model"
	"comment/rpc/types/comment"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	commentInsert := new(model.Comment)
	commentInsert.UserId = 1
	commentInsert.VideoId = 1
	var StatusCode int32 = 0
	var statusMsg = "测试"
	if _, err := l.svcCtx.CommentModel.Insert(l.ctx, commentInsert); err != nil {
		fmt.Printf("插入失败")
		StatusCode = 2
		statusMsg = "测试失败"
		return &comment.DouyinCommentActionResponse{
			StatusCode: &StatusCode,
			StatusMsg:  &statusMsg,
		}, nil
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: &StatusCode,
		StatusMsg:  &statusMsg,
	}, nil
}
