package main

import (
	"flag"
	c2 "github.com/Guo-Chenxu/go-zero-nacos/configcenter"
	"github.com/Guo-Chenxu/go-zero-nacos/register"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"nacos-demo/app/user/cmd/rpc/internal/config"
	"nacos-demo/app/user/cmd/rpc/internal/server"
	"nacos-demo/app/user/cmd/rpc/internal/svc"
	"nacos-demo/app/user/cmd/rpc/pb"
)

var bootstrapConfigFile = flag.String("f", "etc/bootstrap.yaml", "the config.go file")

func main() {
	//解析bootstrap并获取业务配置
	flag.Parse()
	var bootstrapConfig c2.BootstrapConfig
	var c config.Config

	c2.GetConfig(bootstrapConfigFile, &bootstrapConfig, &c)

	// 进行业务配置
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserRpcServer(grpcServer, server.NewUserRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 注册nacos
	register.Register(c)

	s.Start()
}
