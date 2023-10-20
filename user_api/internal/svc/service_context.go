package svc

import (
	"log"
	"user_api/internal/config"
	"user_api/internal/middleware"
	"user_api/model"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleWare rest.Middleware
	UserModel      model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:         c,
		TestMiddleWare: middleware.NewTestMiddleWareMiddleware().Handle,
		UserModel:      model.NewUserModel(db, c.Cache),
	}
}
