package userOpt

import (
	"context"
	"douyin/pkg/comment/common/xerr"
	"douyin/pkg/comment/rpc/usercomment"

	"douyin/pkg/comment/api/internal/svc"
	"douyin/pkg/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (resp *types.CommentListRes, err error) {
	//获取参数
	// 调用rpc 更新user_comment表
	_, err = l.svcCtx.UserCommentRpc.GetVideoComment(l.ctx, &usercomment.GetVideoCommentReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		logx.Errorf("UserCommentList->commentRpc  err : %v , val : %s , message:%+v", err)
		return &types.CommentListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to commentRpc err",
			},
		}, nil
	}

	return &types.CommentListRes{
		Status: types.Status{
			Code: xerr.OK,
		},
	}, nil

}
