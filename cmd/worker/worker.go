package worker

import (
	"fmt"
	"net/http"

	"github.com/jeremykane/go-boilerplate/internal/app"
	"github.com/jeremykane/go-boilerplate/internal/config"
)

func Start(cfg *config.Config) {
	fmt.Println("Starting dealer worker...")
	httpServer := http.Server{}
	app.RunWorker(&httpServer, cfg)
}
