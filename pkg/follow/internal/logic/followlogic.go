package logic

import (
	"context"
	"database/sql"
	"douyin/common/model/followModel"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"
	"douyin/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *follow.FollowReq) (*follow.FollowResp, error) {

	followId, err := l.svcCtx.FollowModel.FindIfExist(l.ctx, in.UserId, in.ToUserId)
	if err != nil {
		logger.Errorf("follow option failed: %v", err.Error())
		return &follow.FollowResp{
			StatusCode: -1,
			StatusMsg:  "follow option failed",
		}, err
	}
	if followId != 0 {
		var removed int64
		if in.ActionType == 1 {
			removed = 1
		}
		// 查询到了followId，调用update更新
		newFollowModel := &followModel.Follow{
			Id:      followId,
			Removed: removed,
		}
		err = l.svcCtx.FollowModel.Update(l.ctx, newFollowModel)
		if err != nil {
			logger.Errorf("follow option failed: %v", err.Error())
			return &follow.FollowResp{
				StatusCode: -1,
				StatusMsg:  "follow option failed",
			}, err
		}
	} else {
		// 未查询到相关数据，直接insert
		if in.ActionType == 1 {
			newFollowModel := &followModel.Follow{
				UserId: sql.NullInt64{
					Int64: in.UserId,
					Valid: true,
				},
				FunId: in.ToUserId,
			}
			_, err = l.svcCtx.FollowModel.Insert(l.ctx, newFollowModel)
			if err != nil {
				return &follow.FollowResp{
					StatusCode: -1,
					StatusMsg:  "follow option failed",
				}, err
			}
		}
	}

	return &follow.FollowResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
