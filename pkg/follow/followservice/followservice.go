// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package followservice

import (
	"context"

	"douyin/pkg/follow/types/follow"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckIsFollowReq    = follow.CheckIsFollowReq
	CheckIsFollowResp   = follow.CheckIsFollowResp
	FollowReq           = follow.FollowReq
	FollowResp          = follow.FollowResp
	FriendUser          = follow.FriendUser
	GetFollowListReq    = follow.GetFollowListReq
	GetFollowListResp   = follow.GetFollowListResp
	GetFollowerListReq  = follow.GetFollowerListReq
	GetFollowerListResp = follow.GetFollowerListResp
	GetFriendListReq    = follow.GetFriendListReq
	GetFriendListResp   = follow.GetFriendListResp
	User                = follow.User

	FollowService interface {
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
		GetFollowList(ctx context.Context, in *GetFollowListReq, opts ...grpc.CallOption) (*GetFollowListResp, error)
		GetFollowerList(ctx context.Context, in *GetFollowerListReq, opts ...grpc.CallOption) (*GetFollowerListResp, error)
		GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error)
		CheckIsFollow(ctx context.Context, in *CheckIsFollowReq, opts ...grpc.CallOption) (*CheckIsFollowResp, error)
	}

	defaultFollowService struct {
		cli zrpc.Client
	}
)

func NewFollowService(cli zrpc.Client) FollowService {
	return &defaultFollowService{
		cli: cli,
	}
}

func (m *defaultFollowService) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultFollowService) GetFollowList(ctx context.Context, in *GetFollowListReq, opts ...grpc.CallOption) (*GetFollowListResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.GetFollowList(ctx, in, opts...)
}

func (m *defaultFollowService) GetFollowerList(ctx context.Context, in *GetFollowerListReq, opts ...grpc.CallOption) (*GetFollowerListResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.GetFollowerList(ctx, in, opts...)
}

func (m *defaultFollowService) GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.GetFriendList(ctx, in, opts...)
}

func (m *defaultFollowService) CheckIsFollow(ctx context.Context, in *CheckIsFollowReq, opts ...grpc.CallOption) (*CheckIsFollowResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.CheckIsFollow(ctx, in, opts...)
}
