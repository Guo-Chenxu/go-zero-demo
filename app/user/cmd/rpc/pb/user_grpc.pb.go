// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user.proto

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

const (
	UserRpc_Ping_FullMethodName = "/pb.User_rpc/Ping"
)

// UserRpcClient is the client API for UserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRpcClient(cc grpc.ClientConnInterface) UserRpcClient {
	return &userRpcClient{cc}
}

func (c *userRpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserRpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRpcServer is the server API for UserRpc service.
// All implementations must embed UnimplementedUserRpcServer
// for forward compatibility
type UserRpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUserRpcServer()
}

// UnimplementedUserRpcServer must be embedded to have forward compatible implementations.
type UnimplementedUserRpcServer struct {
}

func (UnimplementedUserRpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUserRpcServer) mustEmbedUnimplementedUserRpcServer() {}

// UnsafeUserRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRpcServer will
// result in compilation errors.
type UnsafeUserRpcServer interface {
	mustEmbedUnimplementedUserRpcServer()
}

func RegisterUserRpcServer(s grpc.ServiceRegistrar, srv UserRpcServer) {
	s.RegisterService(&UserRpc_ServiceDesc, srv)
}

func _UserRpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRpc_ServiceDesc is the grpc.ServiceDesc for UserRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User_rpc",
	HandlerType: (*UserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UserRpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
