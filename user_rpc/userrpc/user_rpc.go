// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userrpc

import (
	"context"

	"user_rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = pb.Request
	Response = pb.Response

	UserRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}