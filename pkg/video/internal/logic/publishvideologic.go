package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/model/videoModel"
	"douyin/pkg/minio-client/types/minio-client"
	"douyin/pkg/video/internal/svc"
	"douyin/pkg/video/types/video"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	parseToken := token.ParseToken{}
	parseResult, err := parseToken.ParseToken(in.Token)
	if err != nil {
		return &video.PublishVideoResp{
			StatusCode: -1,
			StatusMsg:  "Token鉴权错误",
		}, err
	}
	var authorId = parseResult.UserId
	uploadReq, err := l.svcCtx.MinioRPC.UploadFile(l.ctx, &minio_client.UploadFileRequest{
		Data:  in.Data,
		Title: in.Title,
	})
	if err != nil {
		return &video.PublishVideoResp{
			StatusCode: -1,
			StatusMsg:  "视频上传失败",
		}, err
	}
	videoURL, frontImgURL := uploadReq.VideoUrl, uploadReq.FrontImgUrl
	publishVideo := &videoModel.Video{
		AuthorId: authorId,
		PlayUrl:  videoURL,
		CoverUrl: frontImgURL,
		Time:     time.Now().Unix(),
		Title:    in.Title,
	}
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, publishVideo)
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
