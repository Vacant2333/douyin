package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"
	"douyin/pkg/user/rpc/userservice"
	"fmt"

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
	fmt.Printf(":::::::::::::::::::::::::::::::::::::::::::::::")
	allCommentInfoData, err := l.svcCtx.UserCommentRpc.GetVideoComment(l.ctx, &usercomment.GetVideoCommentReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		logx.Errorf("UserCommentList->commentRpc  err : %v , val : %s , message:%+v", err)
		return &types.CommentListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to commentRpc err",
			},
		}, err
	}
	var commentList []*types.Comment
	for _, v := range allCommentInfoData.CommentList {
		var comment types.Comment
		comment.CommentId = v.CommentId
		comment.Content = v.Content
		comment.CreateTime = v.CreateDate
		userInfoResult, err := l.svcCtx.UserRpc.Info(l.ctx, &userservice.UserInfoReq{
			UserId: v.UserId,
		})
		if err != nil {
			logx.Errorf("UserCommentList->userInfoRpc  err : %v , val : %s , message:%+v", err)
			return &types.CommentListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "send message to userInfoRpc err",
				},
			}, nil
		}

		var user = types.User{
			UserId:        userInfoResult.User.UserId,
			UserName:      userInfoResult.User.UserName,
			FollowCount:   userInfoResult.User.FollowCount,
			FollowerCount: userInfoResult.User.FollowerCount,
			IsFollow:      false,
		}
		comment.User = user

		commentList = append(commentList, &comment)
	}

	return &types.CommentListRes{
		Status: types.Status{
			Code: xerr.OK,
			Msg:  "get comment list success",
		},
		CommentList: commentList,
	}, nil

}
