// Package router 路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/go-backend-tmpl/internal/handler"
)

// RegisterRouter 注册路由
//
//	param app *fiber.App
//	author centonhuang
//	update 2025-01-04 15:32:40
func RegisterRouter(app *fiber.App) fiber.Router {
	pingService := handler.NewPingHandler()

	rootRouter := app.Group("/")
	rootRouter.Get("", pingService.HandlePing)
	v1Router := rootRouter.Group("/v1")
	{
		initTokenRouter(v1Router)
		initOauth2Router(v1Router)
		initUserRouter(v1Router)
	}

	return rootRouter
}
