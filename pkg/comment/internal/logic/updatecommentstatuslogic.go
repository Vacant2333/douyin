package logic

import (
	"context"
	"douyin/common/globalkey"
	"douyin/common/messageTypes"
	"douyin/common/model/commentModel"
	"douyin/common/xerr"
	"douyin/pkg/comment/internal/svc"
	"douyin/pkg/comment/userCommentPb"
	"douyin/pkg/video/videoservice"
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
	newCommentModel := new(commentModel.Comment)

	switch in.ActionType {
	//新增评论
	case messageTypes.ActionADD:
		// Todo 更新返回字段
		newCommentModel.UserId = in.UserId
		newCommentModel.VideoId = in.VideoId
		newCommentModel.Content = in.Content
		insertCommentResult, err := l.svcCtx.UserCommentModel.Insert(l.ctx, newCommentModel)

		if err != nil {
			logx.Errorf("UpdateCommentStatus------->Insert err : %s", err.Error())
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		insertCommentId, err := insertCommentResult.LastInsertId()
		if err != nil {
			logx.Error("UpdateCommentStatus-------> trans fail")
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		_, err = l.svcCtx.VideoRPC.ChangeVideoComment(l.ctx, &videoservice.ChangeVideoCommentReq{
			VideoId:    in.VideoId,
			ActionType: in.ActionType,
		})
		if err != nil {
			logx.Errorf("ChangeVideoComment failed %s ", err)
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		return &userCommentPb.UpdateCommentStatusResp{
			CommentId: insertCommentId,
		}, nil

	//删除评论
	case messageTypes.ActionCancel:
		newCommentModel.Removed = globalkey.DelStateYes
		newCommentModel.Id = in.CommentId
		newCommentModel.VideoId = in.VideoId
		newCommentModel.UserId = in.UserId
		err := l.svcCtx.UserCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			err := l.svcCtx.UserCommentModel.Update(l.ctx, newCommentModel)
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

		_, err = l.svcCtx.VideoRPC.ChangeVideoComment(l.ctx, &videoservice.ChangeVideoCommentReq{
			VideoId:    in.VideoId,
			ActionType: in.ActionType,
		})
		if err != nil {
			logx.Errorf("ChangeVideoComment failed %s ", err)
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

		return &userCommentPb.UpdateCommentStatusResp{
			CommentId: in.CommentId,
		}, nil

	default:
		return &userCommentPb.UpdateCommentStatusResp{}, xerr.NewErrMsg("actionType must be 1 or 2")
	}
}
