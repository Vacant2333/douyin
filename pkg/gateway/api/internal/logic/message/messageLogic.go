package message

import (
	"context"

	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageLogic) Message(req *types.MessageReq) (resp *types.MessageRes, err error) {
	// todo: add your logic here and delete this line

	return
}
