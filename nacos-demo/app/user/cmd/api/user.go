package main

import (
	"flag"
	"fmt"
	c2 "github.com/Guo-Chenxu/go-zero-nacos/configcenter"
	"nacos-demo/app/user/cmd/api/internal/config"
	"nacos-demo/app/user/cmd/api/internal/handler"
	"nacos-demo/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

var bootstrapConfigFile = flag.String("f", "etc/bootstrap.yaml", "the config.go file")

func main() {
	//解析bootstrap config.go
	flag.Parse()
	var bootstrapConfig c2.BootstrapConfig
	var c config.Config
	c2.GetConfig(bootstrapConfigFile, &bootstrapConfig, &c)

	// 配置业务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
