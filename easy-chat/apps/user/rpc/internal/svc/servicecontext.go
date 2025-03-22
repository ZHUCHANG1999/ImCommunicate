package svc

import (
	models2 "easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	models2.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UsersModel: models2.NewUsersModel(sqlConn, c.Cache),
	}
}
