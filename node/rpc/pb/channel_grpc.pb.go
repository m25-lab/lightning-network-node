// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: node/rpc/proto/channel.proto

package pb

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

// ChannelServiceClient is the client API for ChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelServiceClient interface {
	OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelResponse, error)
	CreateCommitment(ctx context.Context, in *CreateCommitmentRequest, opts ...grpc.CallOption) (*CreateCommitmentResponse, error)
	WithdrawHashlock(ctx context.Context, in *WithdrawHashlockRequest, opts ...grpc.CallOption) (*WithdrawHashlockResponse, error)
	WithdrawTimelock(ctx context.Context, in *WithdrawTimelockRequest, opts ...grpc.CallOption) (*WithdrawTimelockResponse, error)
}

type channelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelServiceClient(cc grpc.ClientConnInterface) ChannelServiceClient {
	return &channelServiceClient{cc}
}

func (c *channelServiceClient) OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelResponse, error) {
	out := new(OpenChannelResponse)
	err := c.cc.Invoke(ctx, "/ChannelService/OpenChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) CreateCommitment(ctx context.Context, in *CreateCommitmentRequest, opts ...grpc.CallOption) (*CreateCommitmentResponse, error) {
	out := new(CreateCommitmentResponse)
	err := c.cc.Invoke(ctx, "/ChannelService/CreateCommitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) WithdrawHashlock(ctx context.Context, in *WithdrawHashlockRequest, opts ...grpc.CallOption) (*WithdrawHashlockResponse, error) {
	out := new(WithdrawHashlockResponse)
	err := c.cc.Invoke(ctx, "/ChannelService/WithdrawHashlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) WithdrawTimelock(ctx context.Context, in *WithdrawTimelockRequest, opts ...grpc.CallOption) (*WithdrawTimelockResponse, error) {
	out := new(WithdrawTimelockResponse)
	err := c.cc.Invoke(ctx, "/ChannelService/WithdrawTimelock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServiceServer is the server API for ChannelService service.
// All implementations must embed UnimplementedChannelServiceServer
// for forward compatibility
type ChannelServiceServer interface {
	OpenChannel(context.Context, *OpenChannelRequest) (*OpenChannelResponse, error)
	CreateCommitment(context.Context, *CreateCommitmentRequest) (*CreateCommitmentResponse, error)
	WithdrawHashlock(context.Context, *WithdrawHashlockRequest) (*WithdrawHashlockResponse, error)
	WithdrawTimelock(context.Context, *WithdrawTimelockRequest) (*WithdrawTimelockResponse, error)
	mustEmbedUnimplementedChannelServiceServer()
}

// UnimplementedChannelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServiceServer struct {
}

func (UnimplementedChannelServiceServer) OpenChannel(context.Context, *OpenChannelRequest) (*OpenChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenChannel not implemented")
}
func (UnimplementedChannelServiceServer) CreateCommitment(context.Context, *CreateCommitmentRequest) (*CreateCommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommitment not implemented")
}
func (UnimplementedChannelServiceServer) WithdrawHashlock(context.Context, *WithdrawHashlockRequest) (*WithdrawHashlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawHashlock not implemented")
}
func (UnimplementedChannelServiceServer) WithdrawTimelock(context.Context, *WithdrawTimelockRequest) (*WithdrawTimelockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawTimelock not implemented")
}
func (UnimplementedChannelServiceServer) mustEmbedUnimplementedChannelServiceServer() {}

// UnsafeChannelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServiceServer will
// result in compilation errors.
type UnsafeChannelServiceServer interface {
	mustEmbedUnimplementedChannelServiceServer()
}

func RegisterChannelServiceServer(s grpc.ServiceRegistrar, srv ChannelServiceServer) {
	s.RegisterService(&ChannelService_ServiceDesc, srv)
}

func _ChannelService_OpenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).OpenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChannelService/OpenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).OpenChannel(ctx, req.(*OpenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_CreateCommitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).CreateCommitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChannelService/CreateCommitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).CreateCommitment(ctx, req.(*CreateCommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_WithdrawHashlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawHashlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).WithdrawHashlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChannelService/WithdrawHashlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).WithdrawHashlock(ctx, req.(*WithdrawHashlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_WithdrawTimelock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawTimelockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).WithdrawTimelock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChannelService/WithdrawTimelock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).WithdrawTimelock(ctx, req.(*WithdrawTimelockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelService_ServiceDesc is the grpc.ServiceDesc for ChannelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChannelService",
	HandlerType: (*ChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenChannel",
			Handler:    _ChannelService_OpenChannel_Handler,
		},
		{
			MethodName: "CreateCommitment",
			Handler:    _ChannelService_CreateCommitment_Handler,
		},
		{
			MethodName: "WithdrawHashlock",
			Handler:    _ChannelService_WithdrawHashlock_Handler,
		},
		{
			MethodName: "WithdrawTimelock",
			Handler:    _ChannelService_WithdrawTimelock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/rpc/proto/channel.proto",
}