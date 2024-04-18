package rest

import (
	"net/http"

	_ "github.com/jeremykane/go-boilerplate/docs"
	"github.com/jeremykane/go-boilerplate/internal/app"
	"github.com/jeremykane/go-boilerplate/internal/config"
)

func Start(cfg *config.Config) {
	httpServer := http.Server{}
	app.RunServer(&httpServer, cfg)
}
