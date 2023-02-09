// Code generated by goctl. DO NOT EDIT.
// Source: userinfo.proto

package userinfoclient

import (
	"context"

	"douyin/pkg/userinfo-demo/rpc/types/userinfo"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	User        = userinfo.User
	UserinfoReq = userinfo.UserinfoReq
	UserinfoRes = userinfo.UserinfoRes

	Userinfo interface {
		GetUser(ctx context.Context, in *UserinfoReq, opts ...grpc.CallOption) (*UserinfoRes, error)
	}

	defaultUserinfo struct {
		cli zrpc.Client
	}
)

func NewUserinfo(cli zrpc.Client) Userinfo {
	return &defaultUserinfo{
		cli: cli,
	}
}

func (m *defaultUserinfo) GetUser(ctx context.Context, in *UserinfoReq, opts ...grpc.CallOption) (*UserinfoRes, error) {
	client := userinfo.NewUserinfoClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}