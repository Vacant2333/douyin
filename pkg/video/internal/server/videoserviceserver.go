// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"
	logic2 "douyin/pkg/video/internal/logic"
	"douyin/pkg/video/internal/svc"
	video2 "douyin/pkg/video/types/video"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	video2.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServiceServer) PublishVideo(ctx context.Context, in *video2.PublishVideoReq) (*video2.PublishVideoResp, error) {
	l := logic2.NewPublishVideoLogic(ctx, s.svcCtx)
	return l.PublishVideo(in)
}

func (s *VideoServiceServer) GetVideo(ctx context.Context, in *video2.GetVideoReq) (*video2.GetVideoResp, error) {
	l := logic2.NewGetVideoLogic(ctx, s.svcCtx)
	return l.GetVideo(in)
}

func (s *VideoServiceServer) GetAllVideoByUserId(ctx context.Context, in *video2.GetAllVideoByUserIdReq) (*video2.GetAllVideoByUserIdResp, error) {
	l := logic2.NewGetAllVideoByUserIdLogic(ctx, s.svcCtx)
	return l.GetAllVideoByUserId(in)
}