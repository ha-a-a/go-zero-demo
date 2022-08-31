package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/book/user/api/internal/config"
	"go-zero-demo/book/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.BookUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewBookUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
