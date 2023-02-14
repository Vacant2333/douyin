package publish

import (
	"context"

	"douyin/pkg/comment/api/internal/svc"
	"douyin/pkg/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublishVideoListLogic) GetPublishVideoList(req *types.GetPubVideoListReq) (resp *types.GetPubVideoListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
