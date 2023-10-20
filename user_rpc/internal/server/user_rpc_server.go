// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"user_rpc/internal/logic"
	"user_rpc/internal/svc"
	"user_rpc/pb"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRpcServer) Ping(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
