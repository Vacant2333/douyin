package publish

import (
	"context"
	"douyin/common/xerr"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"douyin/pkg/video/types/video"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PubVideoReq) (resp *types.PubVideoRes, err error) {
	pubResp, err := l.svcCtx.VideoRPC.PublishVideo(l.ctx, &video.PublishVideoReq{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	})
	if err != nil {
		logx.Errorf("publish video failed: %v", err.Error())
		return &types.PubVideoRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get videos failed",
			},
		}, err
	}
	return &types.PubVideoRes{
		Status: types.Status{
			Code: int64(pubResp.StatusCode),
			Msg:  pubResp.StatusMsg,
		},
	}, nil
}
