package logic

import (
	"context"
	"fmt"

	"comment/rpc/internal/svc"
	"comment/rpc/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	fmt.Printf("rpc-list参数：%s\\n", in)
	var id int64 = 1
	_, _ = l.svcCtx.CommentModel.FindOne(l.ctx, id)
	var StatusCode int32 = 0
	var statusMsg = "测试"

	return &comment.DouyinCommentListResponse{
		StatusCode: &StatusCode,
		StatusMsg:  &statusMsg,
	}, nil
}
