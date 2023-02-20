package logic

import (
	"context"
	"douyin/common/globalkey"
	"douyin/common/messageTypes"
	"douyin/common/model/commentModel"
	"douyin/common/xerr"
	"douyin/pkg/comment/internal/svc"
	"douyin/pkg/comment/userCommentPb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentStatusLogic {
	return &UpdateCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCommentStatus -----------------------userCommentStatus-----------------------
func (l *UpdateCommentStatusLogic) UpdateCommentStatus(in *userCommentPb.UpdateCommentStatusReq) (*userCommentPb.UpdateCommentStatusResp, error) {
	commentModel := new(commentModel.Comment)

	switch in.ActionType {
	//新增评论
	case messageTypes.ActionADD:
		// Todo 更新返回字段
		commentModel.UserId = in.UserId
		commentModel.VideoId = in.VideoId
		commentModel.Content = in.Content
		insertCommentResult, err := l.svcCtx.UserCommentModel.Insert(l.ctx, commentModel)

		if err != nil {
			logx.Errorf("UpdateCommentStatus------->Insert err : %s", err.Error())
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		insertCommentId, err := insertCommentResult.LastInsertId()
		if err != nil {
			logx.Error("UpdateCommentStatus-------> trans fail")
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		return &userCommentPb.UpdateCommentStatusResp{
			CommentId: insertCommentId,
		}, nil

	//删除评论
	case messageTypes.ActionCancel:
		commentModel.Removed = globalkey.DelStateYes
		commentModel.Id = in.CommentId
		commentModel.VideoId = in.VideoId
		commentModel.UserId = in.UserId
		err := l.svcCtx.UserCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			err := l.svcCtx.UserCommentModel.Update(l.ctx, commentModel)
			if err != nil {
				logx.Errorf("UpdateCommentStatus------->update err : %s", err.Error())
				return err
			}
			return nil
		})

		if err != nil {
			logx.Error("UpdateCommentStatus-------> trans fail")
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}
		return &userCommentPb.UpdateCommentStatusResp{
			CommentId: in.CommentId,
		}, nil

	default:
		return &userCommentPb.UpdateCommentStatusResp{}, xerr.NewErrMsg("actionType must be 1 or 2")
	}
}
