package api

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/config"
)

// NewFiberApp 创建 Fiber 应用实例。
func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		Prefork:                 false,
		ReadTimeout:             config.ReadTimeout,
		WriteTimeout:            config.WriteTimeout,
		IdleTimeout:             constant.IdleTimeout,
		JSONEncoder:             sonic.Marshal,
		JSONDecoder:             sonic.Unmarshal,
		EnableTrustedProxyCheck: true,
		TrustedProxies:          config.TrustedProxies,
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})
}
