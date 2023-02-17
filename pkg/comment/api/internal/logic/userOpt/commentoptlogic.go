package userOpt

import (
	"context"
	"douyin/pkg/comment/api/internal/types"
	"douyin/pkg/comment/common/help/sensitiveWords"
	"douyin/pkg/comment/common/xerr"
	"douyin/pkg/comment/rpc/usercomment"
	"douyin/pkg/userinfo-demo/rpc/types/userinfo"
	"time"

	"douyin/pkg/comment/api/internal/svc"

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
	//敏感词过滤
	safeContent := sensitiveWords.SensitiveWordsFliter(sensitiveWords.SensitiveWords, req.CommentText, '?')

	// 调用rpc 更新user_comment表
	userCommentResult, err := l.svcCtx.UserCommentRpc.UpdateCommentStatus(l.ctx, &usercomment.UpdateCommentStatusReq{
		VideoId:    req.VideoId,
		UserId:     1, /*l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)*/
		Content:    safeContent,
		CommentId:  req.CommentId,
		ActionType: req.ActionType,
	})
	if err != nil {
		logx.Errorf("UserCommentOpt->UserCommentRpc  err : %v , val : %s , message:%+v", err)
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to UserCommentRpc err",
			},
		}, nil
	}

	if req.ActionType == 1 {
		// 调用user-info rpc拉取发布消息的用户信息
		userInfoResult, err := l.svcCtx.UserInfoRpc.GetUser(l.ctx, &userinfo.UserinfoRequest{UserId: "1"})
		if err != nil {
			logx.Errorf("UserCommentOpt->userInfoRpc  err : %v , val : %s , message:%+v", err)
			return &types.CommentOptRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "send message to userInfoRpc err",
				},
			}, nil
		}

		var user = types.User{
			UserId:        1,
			UserName:      userInfoResult.User.Name,
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
				CommentId:  userCommentResult.CommentId,
				Content:    safeContent,
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
