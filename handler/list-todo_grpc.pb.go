// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: list-todo.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ListToDoClient is the client API for ListToDo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListToDoClient interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type listToDoClient struct {
	cc grpc.ClientConnInterface
}

func NewListToDoClient(cc grpc.ClientConnInterface) ListToDoClient {
	return &listToDoClient{cc}
}

func (c *listToDoClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/handler.ListToDo/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListToDoServer is the server API for ListToDo service.
// All implementations must embed UnimplementedListToDoServer
// for forward compatibility
type ListToDoServer interface {
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	mustEmbedUnimplementedListToDoServer()
}

// UnimplementedListToDoServer must be embedded to have forward compatible implementations.
type UnimplementedListToDoServer struct {
}

func (UnimplementedListToDoServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedListToDoServer) mustEmbedUnimplementedListToDoServer() {}

// UnsafeListToDoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListToDoServer will
// result in compilation errors.
type UnsafeListToDoServer interface {
	mustEmbedUnimplementedListToDoServer()
}

func RegisterListToDoServer(s grpc.ServiceRegistrar, srv ListToDoServer) {
	s.RegisterService(&ListToDo_ServiceDesc, srv)
}

func _ListToDo_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListToDoServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/handler.ListToDo/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListToDoServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ListToDo_ServiceDesc is the grpc.ServiceDesc for ListToDo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListToDo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "handler.ListToDo",
	HandlerType: (*ListToDoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ListToDo_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "list-todo.proto",
}
