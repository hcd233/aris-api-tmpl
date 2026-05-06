package bootstrap

import (
	"github.com/hcd233/aris-api-tmpl/internal/config"
	"github.com/hcd233/aris-api-tmpl/internal/enum"
	"github.com/hcd233/aris-api-tmpl/internal/handler"
	"github.com/hcd233/aris-api-tmpl/internal/router"
	"go.uber.org/dig"
)

type routeParams struct {
	dig.In

	PingHandler   handler.PingHandler
	TokenHandler  handler.TokenHandler
	Oauth2Handler handler.Oauth2Handler
	UserHandler   handler.UserHandler
}

// RegisterRoutes 注册文档和 API 路由。
func RegisterRoutes(server *Server) error {
	return server.container.Invoke(func(params routeParams) {
		if config.Env != enum.EnvProduction {
			router.RegisterDocsRouter(server.App)
		}
		router.RegisterAPIRouter(server.HumaAPI, router.APIRouterDependencies{
			PingHandler:   params.PingHandler,
			TokenHandler:  params.TokenHandler,
			Oauth2Handler: params.Oauth2Handler,
			UserHandler:   params.UserHandler,
		})
	})
}
