// Code generated by goctl. DO NOT EDIT!
// Source: minio-client.proto

package server

import (
	"context"

	"douyin/pkg/minio-client/internal/logic"
	"douyin/pkg/minio-client/internal/svc"
	"douyin/pkg/minio-client/types/minio-client"
)

type UploadFileServer struct {
	svcCtx *svc.ServiceContext
	minio_client.UnimplementedUploadFileServer
}

func NewUploadFileServer(svcCtx *svc.ServiceContext) *UploadFileServer {
	return &UploadFileServer{
		svcCtx: svcCtx,
	}
}

func (s *UploadFileServer) UploadFile(ctx context.Context, in *minio_client.UploadFileRequest) (*minio_client.UploadFileReply, error) {
	l := logic.NewUploadFileLogic(ctx, s.svcCtx)
	return l.UploadFile(in)
}
