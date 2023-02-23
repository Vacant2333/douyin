// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package server

import (
	"context"

	"douyin/pkg/follow/internal/logic"
	"douyin/pkg/follow/internal/svc"
	"douyin/pkg/follow/types/follow"
)

type FollowServiceServer struct {
	svcCtx *svc.ServiceContext
	follow.UnimplementedFollowServiceServer
}

func NewFollowServiceServer(svcCtx *svc.ServiceContext) *FollowServiceServer {
	return &FollowServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *FollowServiceServer) Follow(ctx context.Context, in *follow.FollowReq) (*follow.FollowResp, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}

func (s *FollowServiceServer) GetFollowList(ctx context.Context, in *follow.GetFollowListReq) (*follow.GetFollowListResp, error) {
	l := logic.NewGetFollowListLogic(ctx, s.svcCtx)
	return l.GetFollowList(in)
}

func (s *FollowServiceServer) GetFollowerList(ctx context.Context, in *follow.GetFollowerListReq) (*follow.GetFollowerListResp, error) {
	l := logic.NewGetFollowerListLogic(ctx, s.svcCtx)
	return l.GetFollowerList(in)
}

func (s *FollowServiceServer) CheckIsFollow(ctx context.Context, in *follow.CheckIsFollowReq) (*follow.CheckIsFollowResp, error) {
	l := logic.NewCheckIsFollowLogic(ctx, s.svcCtx)
	return l.CheckIsFollow(in)
}