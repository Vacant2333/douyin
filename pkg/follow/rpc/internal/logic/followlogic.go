package logic

import (
	"context"
	"fmt"

	"douyin/pkg/follow/rpc/internal/svc"
	"douyin/pkg/follow/rpc/types/follow"

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
	// todo: 验证token，取出用户id
	fmt.Println(in.Token)
	userid := 1
	// todo: 将操作加入到消息队列，如果错误回报错
	fmt.Println(in.ActionType, in.ToUserId, userid)

	var err error = nil
	if err != nil {
		return nil, err
	}
	return &follow.FollowResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
