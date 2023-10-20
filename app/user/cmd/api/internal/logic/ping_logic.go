package logic

import (
	"context"
	"nacos-demo/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"nacos-demo/app/user/cmd/api/internal/svc"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp string, err error) {
	// todo: add your logic here and delete this line
	ping, err := l.svcCtx.UserRpcClient.Ping(l.ctx, &pb.Request{})
	if err != nil {
		return err.Error(), err
	}
	return ping.Pong, err
}
