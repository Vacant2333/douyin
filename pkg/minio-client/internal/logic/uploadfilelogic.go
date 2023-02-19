package logic

import (
	"bytes"
	"context"
	"douyin/pkg/logger"
	"douyin/pkg/minio-client/internal/svc"
	"douyin/pkg/minio-client/types/minio-client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) UploadFile(in *minio_client.UploadFileRequest) (*minio_client.UploadFileReply, error) {
	pngFrame, err := getVideoFrame(in.Data, 1)
	if err != nil {
		logger.Fatalf("Fail to get video frame, err: %v", err)
		return nil, err
	}

	client := makeMinIOClient()
	bucket := "douyin"
	if in.Data == nil || in.Title == "" {
		logger.Fatal("UploadFile's parameter cant be nil")
		panic("UploadFile's parameter cant be nil")
	}

	videoUrl, err := uploadFile(client, bytes.NewReader(in.Data), in.Title+".mp4", bucket, "")
	if err != nil {
		logger.Fatalf("Fail to upload video file, err: %v", err)
		return nil, err
	}
	frameUrl, err := uploadFile(client, pngFrame, in.Title+".png", bucket, "")
	if err != nil {
		logger.Fatalf("Fail to upload frame file, err: %v, err")
		return nil, err
	}

	return &minio_client.UploadFileReply{
		Success:     true,
		VideoUrl:    videoUrl,
		FrontImgUrl: frameUrl,
	}, nil
}
