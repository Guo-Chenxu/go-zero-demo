package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"nacos-demo/app/user/cmd/api/internal/config"
	"nacos-demo/app/user/cmd/rpc/pb"
	"nacos-demo/app/user/cmd/rpc/userrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient pb.UserRpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConfig)),
	}
}
