// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/app/plugins/grpc/reserve.proto

package grpc

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

// ReserveClient is the client API for Reserve service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReserveClient interface {
	Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error)
}

type reserveClient struct {
	cc grpc.ClientConnInterface
}

func NewReserveClient(cc grpc.ClientConnInterface) ReserveClient {
	return &reserveClient{cc}
}

func (c *reserveClient) Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error) {
	out := new(ReserveResponse)
	err := c.cc.Invoke(ctx, "/Reserve/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReserveServer is the server API for Reserve service.
// All implementations must embed UnimplementedReserveServer
// for forward compatibility
type ReserveServer interface {
	Reserve(context.Context, *ReserveRequest) (*ReserveResponse, error)
}

// UnimplementedReserveServer must be embedded to have forward compatible implementations.
type UnimplementedReserveServer struct {
}

func (UnimplementedReserveServer) Reserve(context.Context, *ReserveRequest) (*ReserveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedReserveServer) mustEmbedUnimplementedReserveServer() {}

// UnsafeReserveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReserveServer will
// result in compilation errors.
type UnsafeReserveServer interface {
	mustEmbedUnimplementedReserveServer()
}

func RegisterReserveServer(s grpc.ServiceRegistrar, srv ReserveServer) {
	s.RegisterService(&Reserve_ServiceDesc, srv)
}

func _Reserve_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReserveServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Reserve/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReserveServer).Reserve(ctx, req.(*ReserveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reserve_ServiceDesc is the grpc.ServiceDesc for Reserve service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reserve_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Reserve",
	HandlerType: (*ReserveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reserve",
			Handler:    _Reserve_Reserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/plugins/grpc/reserve.proto",
}