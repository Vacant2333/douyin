// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: UserMessage.proto

package userMessagePb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserMessageClient is the client API for UserMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMessageClient interface {
	// -----------------------SendMessage-----------------------
	SendMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error)
	// -----------------------GetMessageList-----------------------
	GetMessageList(ctx context.Context, in *MessageListReq, opts ...grpc.CallOption) (*MessageListRes, error)
}

type userMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMessageClient(cc grpc.ClientConnInterface) UserMessageClient {
	return &userMessageClient{cc}
}

func (c *userMessageClient) SendMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error) {
	out := new(MessageRes)
	err := c.cc.Invoke(ctx, "/pb.UserMessage/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMessageClient) GetMessageList(ctx context.Context, in *MessageListReq, opts ...grpc.CallOption) (*MessageListRes, error) {
	out := new(MessageListRes)
	err := c.cc.Invoke(ctx, "/pb.UserMessage/GetMessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMessageServer is the server API for UserMessage service.
// All implementations must embed UnimplementedUserMessageServer
// for forward compatibility
type UserMessageServer interface {
	// -----------------------SendMessage-----------------------
	SendMessage(context.Context, *MessageReq) (*MessageRes, error)
	// -----------------------GetMessageList-----------------------
	GetMessageList(context.Context, *MessageListReq) (*MessageListRes, error)
	mustEmbedUnimplementedUserMessageServer()
}

// UnimplementedUserMessageServer must be embedded to have forward compatible implementations.
type UnimplementedUserMessageServer struct {
}

func (UnimplementedUserMessageServer) SendMessage(context.Context, *MessageReq) (*MessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedUserMessageServer) GetMessageList(context.Context, *MessageListReq) (*MessageListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageList not implemented")
}
func (UnimplementedUserMessageServer) mustEmbedUnimplementedUserMessageServer() {}

// UnsafeUserMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMessageServer will
// result in compilation errors.
type UnsafeUserMessageServer interface {
	mustEmbedUnimplementedUserMessageServer()
}

func RegisterUserMessageServer(s grpc.ServiceRegistrar, srv UserMessageServer) {
	s.RegisterService(&UserMessage_ServiceDesc, srv)
}

func _UserMessage_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserMessage/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageServer).SendMessage(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMessage_GetMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMessageServer).GetMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserMessage/GetMessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMessageServer).GetMessageList(ctx, req.(*MessageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMessage_ServiceDesc is the grpc.ServiceDesc for UserMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserMessage",
	HandlerType: (*UserMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _UserMessage_SendMessage_Handler,
		},
		{
			MethodName: "GetMessageList",
			Handler:    _UserMessage_GetMessageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserMessage.proto",
}
