package logic

import (
	"context"
	"douyin/pkg/comment/common/globalkey"
	"douyin/pkg/comment/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"douyin/pkg/comment/rpc/internal/svc"
	"douyin/pkg/comment/rpc/userCommentPb"

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

// -----------------------userCommentStatus-----------------------
func (l *UpdateCommentStatusLogic) UpdateCommentStatus(in *userCommentPb.UpdateCommentStatusReq) (*userCommentPb.UpdateCommentStatusResp, error) {
	commentModel := new(model.Comment)
	//新增评论
	if in.ActionType == 1 {
		commentModel.UserId = in.UserId
		commentModel.VideoId = in.VideoId
		commentModel.Content = "测试"
		err := l.svcCtx.UserCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {

			// 这里有点不一样 要是model文件里的内容变了 这里也要变
			// InsertOrUpdate(ctx context.Context, session sqlx.Session, field string, setStatus string, videoId, objId, userId, opt int64)
			_, err := l.svcCtx.UserCommentModel.Insert(l.ctx, commentModel)
			if err != nil {
				logx.Errorf("UpdateCommentStatus------->Insert err : %s", err.Error())
				return err
			}
			return nil
		})

		if err != nil {
			logx.Error("UpdateCommentStatus-------> trans fail")
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}
	} else {
		commentModel.Deleted = globalkey.DelStateYes
		err := l.svcCtx.UserCommentModel.Update(l.ctx, commentModel)
		if err != nil {
			logx.Errorf("UpdateCommentStatus------->Insert err : %s", err.Error())
			return &userCommentPb.UpdateCommentStatusResp{}, err
		}

	}
	return &userCommentPb.UpdateCommentStatusResp{}, nil
}
