package router

import (
    "context"
    "net/http"

    "github.com/danielgtaylor/huma/v2"
    "github.com/danielgtaylor/huma/v2/adapters/humafiber"
    "github.com/gofiber/fiber/v2"
    "github.com/hcd233/go-backend-tmpl/internal/protocol"
)

// RegisterHuma 在 Fiber 应用中初始化 Huma，并提供 OpenAPI 文档与示例接口
//
// param app *fiber.App
// author centonhuang
// update 2025-10-29 00:00:00
func RegisterHuma(app *fiber.App) huma.API {
    api := humafiber.New(app, huma.Config{
        OpenAPI: &huma.OpenAPI{
            Info: &huma.Info{
                Title:       "Go Backend API",
                Description: "API 文档（基于 Huma 生成 OpenAPI 规范）",
                Version:     "1.0.0",
            },
        },
    })

    // OpenAPI JSON 规范输出
    app.Get("/openapi.json", func(c *fiber.Ctx) error {
        c.Type("json")
        return c.JSON(api.OpenAPI())
    })

    // 简易在线文档页面（基于 Scalar CDN）
    app.Get("/docs", func(c *fiber.Ctx) error {
        c.Type("html")
        return c.SendString(`<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>API Docs</title>
    <style>html,body{height:100%;margin:0;} body{display:flex;}</style>
  </head>
  <body>
    <script id="api-reference" data-url="/openapi.json" src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`) 
    })

    // 通用响应体（用于让 OpenAPI 包含明确的 data 结构）
    type APIResponse[T any] struct {
        Data  T      `json:"data"`
        Error string `json:"error,omitempty"`
    }

    // 示例：使用 Huma 定义一个 Ping 接口，返回与现有健康检查一致的数据结构
    type PingOutput struct {
        Body APIResponse[protocol.PingResponse]
    }

    huma.Register(api, huma.Operation{
        Method:  http.MethodGet,
        Path:    "/huma/ping",
        Summary: "健康检查（Huma 示例）",
        Tags:    []string{"ping"},
    }, func(_ context.Context, _ *struct{}) (*PingOutput, error) {
        return &PingOutput{
            Body: APIResponse[protocol.PingResponse]{
                Data: protocol.PingResponse{Status: "ok"},
            },
        }, nil
    })

    return api
}
