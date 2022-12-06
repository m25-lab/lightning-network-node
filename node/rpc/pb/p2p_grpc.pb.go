// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: node/rpc/proto/p2p.proto

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

// PeerToPeerClient is the client API for PeerToPeer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerToPeerClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	GetListPeer(ctx context.Context, in *GetListPeerRequest, opts ...grpc.CallOption) (*GetListPeerResponse, error)
	GetListChannel(ctx context.Context, in *GetListChannelRequest, opts ...grpc.CallOption) (*GetListPeerResponse, error)
}

type peerToPeerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerToPeerClient(cc grpc.ClientConnInterface) PeerToPeerClient {
	return &peerToPeerClient{cc}
}

func (c *peerToPeerClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/p2p.PeerToPeer/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerToPeerClient) GetListPeer(ctx context.Context, in *GetListPeerRequest, opts ...grpc.CallOption) (*GetListPeerResponse, error) {
	out := new(GetListPeerResponse)
	err := c.cc.Invoke(ctx, "/p2p.PeerToPeer/GetListPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerToPeerClient) GetListChannel(ctx context.Context, in *GetListChannelRequest, opts ...grpc.CallOption) (*GetListPeerResponse, error) {
	out := new(GetListPeerResponse)
	err := c.cc.Invoke(ctx, "/p2p.PeerToPeer/GetListChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerToPeerServer is the server API for PeerToPeer service.
// All implementations must embed UnimplementedPeerToPeerServer
// for forward compatibility
type PeerToPeerServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	GetListPeer(context.Context, *GetListPeerRequest) (*GetListPeerResponse, error)
	GetListChannel(context.Context, *GetListChannelRequest) (*GetListPeerResponse, error)
	mustEmbedUnimplementedPeerToPeerServer()
}

// UnimplementedPeerToPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerToPeerServer struct {
}

func (UnimplementedPeerToPeerServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedPeerToPeerServer) GetListPeer(context.Context, *GetListPeerRequest) (*GetListPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListPeer not implemented")
}
func (UnimplementedPeerToPeerServer) GetListChannel(context.Context, *GetListChannelRequest) (*GetListPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListChannel not implemented")
}
func (UnimplementedPeerToPeerServer) mustEmbedUnimplementedPeerToPeerServer() {}

// UnsafePeerToPeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerToPeerServer will
// result in compilation errors.
type UnsafePeerToPeerServer interface {
	mustEmbedUnimplementedPeerToPeerServer()
}

func RegisterPeerToPeerServer(s grpc.ServiceRegistrar, srv PeerToPeerServer) {
	s.RegisterService(&PeerToPeer_ServiceDesc, srv)
}

func _PeerToPeer_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerToPeerServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.PeerToPeer/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerToPeerServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerToPeer_GetListPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerToPeerServer).GetListPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.PeerToPeer/GetListPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerToPeerServer).GetListPeer(ctx, req.(*GetListPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerToPeer_GetListChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerToPeerServer).GetListChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.PeerToPeer/GetListChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerToPeerServer).GetListChannel(ctx, req.(*GetListChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerToPeer_ServiceDesc is the grpc.ServiceDesc for PeerToPeer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerToPeer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.PeerToPeer",
	HandlerType: (*PeerToPeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _PeerToPeer_Connect_Handler,
		},
		{
			MethodName: "GetListPeer",
			Handler:    _PeerToPeer_GetListPeer_Handler,
		},
		{
			MethodName: "GetListChannel",
			Handler:    _PeerToPeer_GetListChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/rpc/proto/p2p.proto",
}
