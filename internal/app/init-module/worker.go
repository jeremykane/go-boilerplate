package init_module

import (
	"fmt"

	"github.com/jeremykane/go-boilerplate/internal/app/worker"
)

type (
	// WorkerWrapper wraps a Worker module with its graceful handler
	WorkerWrapper struct {
		Cron            *worker.Worker
		GracefulHandler *APIGracefulHandler
	}

	WorkerGracefulHandler struct{}
)

// NewWorkerWrapper creates a new Worker wrapper instance.
func NewWorkerWrapper(cron *worker.Worker, gracefulHandler *APIGracefulHandler) *WorkerWrapper {
	return &WorkerWrapper{
		Cron:            cron,
		GracefulHandler: gracefulHandler,
	}
}

func NewWorkerGracefulHandler() *WorkerGracefulHandler {
	return &WorkerGracefulHandler{}
}

func (h *WorkerGracefulHandler) ReleaseResource() {
	fmt.Println("Releasing worker resources...")
	defer fmt.Println("Worker resources are gracefully released")
}
