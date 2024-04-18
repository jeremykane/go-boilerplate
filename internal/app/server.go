package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/jeremykane/go-boilerplate/docs"
	"github.com/jeremykane/go-boilerplate/internal/app/api"
	init_module "github.com/jeremykane/go-boilerplate/internal/app/init-module"
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/pkg/grace"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/jeremykane/go-boilerplate/pkg/recovery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer(httpServer *http.Server, cfg *config.Config) {
	ctx := context.Background()
	httpClient := http.Client{
		Timeout: time.Duration(cfg.Server.GlobalTimeout) * time.Millisecond,
	}

	apiWrapper := init_module.NewAPI(ctx, cfg, &httpClient)
	defer apiWrapper.GracefulHandler.ReleaseResource()

	echoServer := echo.New()
	logConfig := middleware.DefaultLoggerConfig
	logConfig.Skipper = func(c echo.Context) bool {
		return strings.Contains(c.Request().URL.Path, "ping") || strings.Contains(c.Request().URL.Path, "metrics")
	}

	echoServer.Use(middleware.LoggerWithConfig(logConfig))
	echoServer.Use(recovery.EchoMiddlewarePanicHandler())
	echoServer.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

	api.RegisterRoutes(echoServer, apiWrapper.Handler, cfg)

	httpServer.Addr = fmt.Sprintf(":%d", cfg.Server.APIPort)
	httpServer.Handler = echoServer
	if err := grace.ServeHTTP(httpServer, httpServer.Addr, 10*time.Second); err != nil {
		logger.Fatalf(ctx, "[FATAL] Failed to serve HTTP server on %s: %s", httpServer.Addr, err.Error())
	}
}
