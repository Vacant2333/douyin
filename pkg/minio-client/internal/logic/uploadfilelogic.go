package logic

import (
	"bytes"
	"context"
	"douyin/pkg/logger"
	"github.com/minio/minio-go/v7"

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
	client := MakeMinIOClient()
	bucket := "douyin"
	if in.Data == nil {
		logger.Fatal("UploadFile's parameter cant be nil")
		panic("UploadFile's parameter cant be nil")
	}
	// Read file's content
	reader := bytes.NewReader(in.Data)
	fileName := "test"
	// Upload file
	_, err := client.PutObject(context.Background(),
		bucket, fileName, reader, reader.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logger.Fatalf("Upload object error: %v", err.Error())
		return &minio_client.UploadFileReply{
			Success: false,
		}, err
	}
	logger.InfoF("Upload file success, fileName: %v len: %v bytes", fileName, reader.Size())
	return &minio_client.UploadFileReply{
		Success: true,
	}, nil
}
