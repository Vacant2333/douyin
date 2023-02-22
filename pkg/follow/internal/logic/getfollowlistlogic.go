package logic

import (
	"context"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"
	"fmt"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *follow.GetFollowListReq) (*follow.GetFollowListResp, error) {
	// todo: 判断token是否合法
	queryFollows, err := l.svcCtx.FollowModel.FindAllByFunId(l.ctx, in.UserId)
	if err != nil {
		return &follow.GetFollowListResp{
			StatusCode: -1,
			StatusMsg:  "Failed, something seems to be wrong on the server side",
		}, err
	}
	userList := make([]*follow.User, len(queryFollows))
	var wg sync.WaitGroup
	for i, queryFollow := range queryFollows {
		query := queryFollow
		wg.Add(1)
		go func(i int) {
			// todo: 生成user model，根据query中的userId查询User数据，并组装UserList
			fmt.Println(query.UserId)
			userList[i].Name = "TBH"
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	return &follow.GetFollowListResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	}, nil
}
