package logic

import (
	"context"
	"douyin/common/model/videoModel"
	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoReq) (*video.PublishVideoResp, error) {
	// todo: 根据token获取authorId
	var authorId int64 = 1
	// todo: 调用minIO RPC发送，获得videoURL,frontImgURL
	videoURL, frontImgURL := "I am videoURL", "I am frontImgURL"

	publishVideo := &videoModel.Video{
		AuthorId: authorId,
		PlayUrl:  videoURL,
		CoverUrl: frontImgURL,
		Time:     time.Now().Unix(),
		Title:    in.Title,
	}
	_, err := l.svcCtx.VideoModel.Insert(l.ctx, publishVideo)
	if err != nil {
		return &video.PublishVideoResp{
			StatusCode: -1,
			StatusMsg:  "Failed, something seems to be wrong on the server side",
		}, err
	}
	return &video.PublishVideoResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
