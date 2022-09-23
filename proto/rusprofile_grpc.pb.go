// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: proto/rusprofile.proto

package proto

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

// RusprofileWrapperClient is the client API for RusprofileWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RusprofileWrapperClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type rusprofileWrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewRusprofileWrapperClient(cc grpc.ClientConnInterface) RusprofileWrapperClient {
	return &rusprofileWrapperClient{cc}
}

func (c *rusprofileWrapperClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RusprofileWrapper/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RusprofileWrapperServer is the server API for RusprofileWrapper service.
// All implementations must embed UnimplementedRusprofileWrapperServer
// for forward compatibility
type RusprofileWrapperServer interface {
	Get(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedRusprofileWrapperServer()
}

// UnimplementedRusprofileWrapperServer must be embedded to have forward compatible implementations.
type UnimplementedRusprofileWrapperServer struct {
}

func (UnimplementedRusprofileWrapperServer) Get(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRusprofileWrapperServer) mustEmbedUnimplementedRusprofileWrapperServer() {}

// UnsafeRusprofileWrapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RusprofileWrapperServer will
// result in compilation errors.
type UnsafeRusprofileWrapperServer interface {
	mustEmbedUnimplementedRusprofileWrapperServer()
}

func RegisterRusprofileWrapperServer(s grpc.ServiceRegistrar, srv RusprofileWrapperServer) {
	s.RegisterService(&RusprofileWrapper_ServiceDesc, srv)
}

func _RusprofileWrapper_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RusprofileWrapperServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RusprofileWrapper/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RusprofileWrapperServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// RusprofileWrapper_ServiceDesc is the grpc.ServiceDesc for RusprofileWrapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RusprofileWrapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RusprofileWrapper",
	HandlerType: (*RusprofileWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RusprofileWrapper_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rusprofile.proto",
}