// Package router 路由。
package router

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/aris-api-tmpl/internal/handler"
)

// APIRouterDependencies API 路由依赖。
type APIRouterDependencies struct {
	PingHandler   handler.PingHandler
	TokenHandler  handler.TokenHandler
	Oauth2Handler handler.Oauth2Handler
	UserHandler   handler.UserHandler
}

// RegisterDocsRouter 注册文档路由。
func RegisterDocsRouter(app *fiber.App) {
	app.Get("/docs", func(c *fiber.Ctx) error {
		html := `<!doctype html>
<html>
  <head>
    <title>Aris API Tmpl Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      data-url="/openapi.json"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`
		return c.Type("html").SendString(html)
	})
}

// RegisterAPIRouter 注册 API 路由。
func RegisterAPIRouter(humaAPI huma.API, deps APIRouterDependencies) {
	apiGroup := huma.NewGroup(humaAPI, "/api")
	v1Group := huma.NewGroup(apiGroup, "/v1")

	initHealthRouter(humaAPI, deps.PingHandler)

	tokenGroup := huma.NewGroup(v1Group, "/token")
	initTokenRouter(tokenGroup, deps.TokenHandler)

	oauth2Group := huma.NewGroup(v1Group, "/oauth2")
	initOauth2Router(oauth2Group, deps.Oauth2Handler)

	userGroup := huma.NewGroup(v1Group, "/user")
	initUserRouter(userGroup, deps.UserHandler)
}
