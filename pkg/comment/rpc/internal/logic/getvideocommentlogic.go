package logic

import (
	"context"
	"douyin/pkg/comment/rpc/internal/svc"
	"douyin/pkg/comment/rpc/userCommentPb"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentLogic {
	return &GetVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetVideoComment -----------------------userCommentList-----------------------
func (l *GetVideoCommentLogic) GetVideoComment(in *userCommentPb.GetVideoCommentReq) (*userCommentPb.GetVideoCommentReqResp, error) {
	fmt.Printf(":::::::::::::::::::::::::::::::::::::::::::::::")
	allCommentInfoData, err := l.svcCtx.UserCommentModel.FindAll(l.ctx, in.VideoId)

	if err != nil {
		logx.Errorf("GetCommentList------->SELECT err : %s", err.Error())
		return &userCommentPb.GetVideoCommentReqResp{}, err
	}

	var commentList []*userCommentPb.Comment
	for _, v := range allCommentInfoData {
		var comment userCommentPb.Comment
		comment.CommentId = v.Id
		comment.Content = v.Content
		comment.CreateDate = v.CreateTime.String()
		comment.UserId = v.UserId

		commentList = append(commentList, &comment)
	}

	return &userCommentPb.GetVideoCommentReqResp{
		CommentList: commentList,
	}, nil
}
