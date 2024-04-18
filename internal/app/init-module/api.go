package init_module

import "github.com/jeremykane/go-boilerplate/internal/app/api"

type (
	// APIWrapper wraps an API module with its graceful handler
	APIWrapper struct {
		Handler         *api.Handler
		GracefulHandler *APIGracefulHandler
	}

	// APIGracefulHandler stores dependencies that need to be released before shutting down the service.
	// Example: kafka
	APIGracefulHandler struct{}
)

// NewAPIWrapper creates a new API wrapper instance.
func NewAPIWrapper(handler *api.Handler, gracefulHandler *APIGracefulHandler) *APIWrapper {
	return &APIWrapper{
		Handler:         handler,
		GracefulHandler: gracefulHandler,
	}
}

// NewAPIGracefulHandler creates a new API graceful handler instance
func NewAPIGracefulHandler() *APIGracefulHandler {
	return &APIGracefulHandler{}
}

// ReleaseResource releases every resource/dependency.
func (h *APIGracefulHandler) ReleaseResource() {
	// Example: flush kafka, wait for queue, etc
}
