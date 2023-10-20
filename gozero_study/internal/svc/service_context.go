package svc

import (
	"gozero_study/internal/config"
	"gozero_study/internal/middleware"
	"gozero_study/model"
	"user_rpc/userrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	TestMiddleWare    rest.Middleware
	MessageMiddleWare rest.Middleware
	UserModel         model.UserModel
	UserRpcClient     userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		TestMiddleWare:    middleware.NewTestMiddleWareMiddleware().Handle,
		MessageMiddleWare: middleware.NewMessageMiddleWareMiddleware().Handle,
		UserModel:         model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserRpcClient:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConfig)),
	}
}
