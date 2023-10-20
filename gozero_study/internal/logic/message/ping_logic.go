package message

import (
	"context"
	"user_rpc/pb"

	"gozero_study/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
	res, err := l.svcCtx.UserRpcClient.Ping(l.ctx, &pb.Request{
		Ping: "ping",
	})
	return res.GetPong(), err
}
