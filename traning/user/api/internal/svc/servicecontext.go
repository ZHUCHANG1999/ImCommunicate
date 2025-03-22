package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"imooc/traning/user/api/internal/config"
	"imooc/traning/user/api/internal/middleware"
	"imooc/traning/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	userclient.User

	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		//User: userclient.NewUser(zrpc.MustNewClient(c.UserPrc)),

		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
