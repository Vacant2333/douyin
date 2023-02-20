package message

import (
	"context"

	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessagelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessagelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessagelistLogic {
	return &MessagelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessagelistLogic) Messagelist(req *types.MessageListReq) (resp *types.MessageListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
