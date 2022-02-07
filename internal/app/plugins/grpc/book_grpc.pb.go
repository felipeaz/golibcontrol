// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: book.proto

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

// GetBookInfoClient is the client API for GetBookInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetBookInfoClient interface {
	GetBookInfo(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type getBookInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewGetBookInfoClient(cc grpc.ClientConnInterface) GetBookInfoClient {
	return &getBookInfoClient{cc}
}

func (c *getBookInfoClient) GetBookInfo(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/GetBookInfo/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetBookInfoServer is the server API for GetBookInfo service.
// All implementations must embed UnimplementedGetBookInfoServer
// for forward compatibility
type GetBookInfoServer interface {
	GetBookInfo(context.Context, *GetBookRequest) (*GetBookResponse, error)
}

// UnimplementedGetBookInfoServer must be embedded to have forward compatible implementations.
type UnimplementedGetBookInfoServer struct {
}

func (UnimplementedGetBookInfoServer) GetBookInfo(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}
func (UnimplementedGetBookInfoServer) mustEmbedUnimplementedGetBookInfoServer() {}

// UnsafeGetBookInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetBookInfoServer will
// result in compilation errors.
type UnsafeGetBookInfoServer interface {
	mustEmbedUnimplementedGetBookInfoServer()
}

func RegisterGetBookInfoServer(s grpc.ServiceRegistrar, srv GetBookInfoServer) {
	s.RegisterService(&GetBookInfo_ServiceDesc, srv)
}

func _GetBookInfo_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetBookInfoServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetBookInfo/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetBookInfoServer).GetBookInfo(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetBookInfo_ServiceDesc is the grpc.ServiceDesc for GetBookInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetBookInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetBookInfo",
	HandlerType: (*GetBookInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _GetBookInfo_GetBookInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
