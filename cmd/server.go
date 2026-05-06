package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/aris-api-tmpl/internal/bootstrap"
	"github.com/hcd233/aris-api-tmpl/internal/config"
	"github.com/hcd233/aris-api-tmpl/internal/cron"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/cache"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/database"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/httpclient"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/pool"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"github.com/hcd233/aris-api-tmpl/internal/middleware"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// shutdownTimeout 优雅关闭的最大超时时间
const shutdownTimeout = 60 * time.Second

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server Command Group",
	Long:  `Server command group for starting and managing the API server`,
}

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the API server",
	Long:  `Start and run the API server, listening on the specified host and port`,
	Run: func(cmd *cobra.Command, _ []string) {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger().Error("[Server] Start server panic", zap.Any("error", r), zap.ByteString("stack", debug.Stack()))
				os.Exit(1)
			}
		}()
		host, port := lo.Must1(cmd.Flags().GetString("host")), lo.Must1(cmd.Flags().GetString("port"))

		logger.Logger().Info("[Server] Environment",
			zap.String("env", config.Env),
			zap.Duration("readTimeout", config.ReadTimeout),
			zap.Duration("writeTimeout", config.WriteTimeout),
			zap.Int("maxHeaderBytes", config.MaxHeaderBytes),
			zap.Int("poolWorkers", config.PoolWorkers),
			zap.Int("poolQueueSize", config.PoolQueueSize),
			zap.Strings("trustedProxies", config.TrustedProxies),
		)

		database.InitDatabase()
		cache.InitCache()
		httpclient.InitHTTPClient()
		pool.InitPoolManager()

		server, err := bootstrap.BuildServer()
		if err != nil {
			logger.Logger().Error("[Server] Build server failed", zap.Error(err))
			os.Exit(1)
		}
		app := server.App
		app.Use(
			middleware.RecoverMiddleware(),
			middleware.FgprofMiddleware(),
			middleware.CORSMiddleware(),
			middleware.CompressMiddleware(),
			middleware.TraceMiddleware(),
			middleware.LogMiddleware(middleware.LogMiddlewareConfig{
				SamplingRules: []middleware.LogSamplingRule{
					{Path: "/health", Interval: 5 * time.Minute},
					{Path: "/ssehealth", Interval: 5 * time.Minute},
				},
			}),
		)
		if err := bootstrap.RegisterRoutes(server); err != nil {
			logger.Logger().Error("[Server] Register routes failed", zap.Error(err))
			os.Exit(1)
		}

		listenAddr := fmt.Sprintf("%s:%s", host, port)
		listenErr := make(chan error, 1)
		go func() {
			listenErr <- app.Listen(listenAddr)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case err := <-listenErr:
			if err != nil {
				logger.Logger().Error("[Server] HTTP server exited unexpectedly", zap.Error(err))
				os.Exit(1)
			}
		case sig := <-quit:
			logger.Logger().Info("[Server] Received shutdown signal, starting graceful shutdown...", zap.String("signal", sig.String()))
			gracefulShutdown(app)
		}
	},
}

// gracefulShutdown 按序执行优雅关闭流程
func gracefulShutdown(app *fiber.App) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)

		logger.Logger().Info("[Server] Step 1/5: Shutting down HTTP server...")
		if err := app.ShutdownWithTimeout(30 * time.Second); err != nil {
			logger.Logger().Error("[Server] HTTP server shutdown error", zap.Error(err))
		}

		logger.Logger().Info("[Server] Step 2/5: Stopping pool manager...")
		pool.StopPoolManager()

		logger.Logger().Info("[Server] Step 3/5: Stopping cron jobs...")
		cron.StopCronJobs()

		logger.Logger().Info("[Server] Step 4/5: Closing database connection...")
		if err := database.CloseDatabase(); err != nil {
			logger.Logger().Error("[Server] Database close error", zap.Error(err))
		}

		logger.Logger().Info("[Server] Step 5/5: Closing Redis connection...")
		if err := cache.CloseCache(); err != nil {
			logger.Logger().Error("[Server] Redis close error", zap.Error(err))
		}

		logger.Logger().Info("[Server] Graceful shutdown completed")
	}()

	select {
	case <-done:
	case <-ctx.Done():
		logger.Logger().Error("[Server] Graceful shutdown timed out, forcing exit", zap.Duration("timeout", shutdownTimeout))
	}
}

func init() {
	serverCmd.AddCommand(startServerCmd)
	rootCmd.AddCommand(serverCmd)

	startServerCmd.Flags().StringP("host", "", "localhost", "监听的主机")
	startServerCmd.Flags().StringP("port", "p", "8080", "监听的端口")
}
