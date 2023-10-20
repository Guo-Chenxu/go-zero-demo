package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	NacosConfig NacosConfig
}

type NacosConfig struct {
	Host                string
	Port                uint64
	Username            string
	Password            string
	NamespaceId         string
	TimeoutMs           uint64
	NotLoadCacheAtStart bool
	LogLevel            string
}
