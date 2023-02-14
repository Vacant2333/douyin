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

// -----------------------userCommentList-----------------------
func (l *GetVideoCommentLogic) GetVideoComment(in *userCommentPb.GetVideoCommentReq) (*userCommentPb.GetVideoCommentReqResp, error) {
	fmt.Printf("GetVideoComment-------------->")
	commentList, err := l.svcCtx.UserCommentModel.FindAll(l.ctx, in.VideoId)
	fmt.Printf("commentList: %v", commentList)
	if err != nil {
		logx.Errorf("GetCommentList------->SELECT err : %s", err.Error())
		return &userCommentPb.GetVideoCommentReqResp{}, err
	}

	return &userCommentPb.GetVideoCommentReqResp{}, nil
}
