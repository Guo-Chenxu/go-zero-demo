package logic

import (
	"context"

	"nacos-demo/app/user/cmd/rpc/internal/svc"
	"nacos-demo/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.Request) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{
		Pong: "pong",
	}, nil
}
