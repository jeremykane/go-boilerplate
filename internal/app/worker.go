package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	init_module "github.com/jeremykane/go-boilerplate/internal/app/init-module"
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/labstack/echo/v4"
)

func RunWorker(httpServer *http.Server, cfg *config.Config) {
	ctx := context.Background()
	httpClient := http.Client{
		Timeout: time.Duration(cfg.Worker.GlobalTimeout) * time.Millisecond,
	}

	worker := init_module.NewWorker(ctx, cfg, &httpClient)
	defer worker.GracefulHandler.ReleaseResource()
	worker.Cron.Start()
	defer worker.Cron.Stop()

	e := echo.New()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Worker.Port),
		Handler: e,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Worker Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Worker Exited")
}
