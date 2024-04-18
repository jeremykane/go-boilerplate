package main

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/jeremykane/go-boilerplate/cmd/worker"
	"github.com/spf13/cobra"

	"github.com/jeremykane/go-boilerplate/cmd/rest"
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/pkg/logger"
	"github.com/jeremykane/go-boilerplate/pkg/logger/logrus"
)

var _, file, _, _ = runtime.Caller(0)
var cfg *config.Config

func main() {
	cmd := &cobra.Command{
		Use:   "go-backend",
		Short: "Add short description here",
		Long:  "",
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:              "server",
			Short:            "HTTP Server Listener",
			Long:             "a http server listener that will be listening http request to this service.",
			TraverseChildren: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				rest.Start(cfg)

				return nil
			},
		},
		&cobra.Command{
			Use:              "worker",
			Short:            "Worker Server",
			Long:             "a worker server that will run all scheduler registered to the worker.",
			TraverseChildren: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				worker.Start(cfg)

				return nil
			},
		})

	cobra.OnInitialize(bootstrap)
	if err := cmd.Execute(); err != nil {
		log.Print(err)
		panic(err)
	}
}

func bootstrap() {
	initConfig()
	initLogger()
}

func initLogger() {
	paramLogs := logrus.LogrusParam{
		Level: cfg.Server.ApiLogLevel,
	}
	logs := logrus.NewLogrus(paramLogs)
	logger.NewLogger(logs)

	logger.InfoWithFields(context.Background(), "initialize logs", logger.Fields{
		"cfg": cfg,
	})
}

func initConfig() {
	var err error
	configPath := filepath.Join(filepath.Dir(file), "../", "config")
	cfg, err = config.Load(configPath, "config")
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}
}
