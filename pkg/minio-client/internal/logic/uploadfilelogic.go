package logic

import (
	"context"

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

	return &minio_client.UploadFileReply{}, nil
}
