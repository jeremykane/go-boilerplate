package grace

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	defaultGraceTimeout = 10 * time.Second

	// ErrGraceShutdownTimeout happens when the server graceful shutdown exceed the given grace timeout
	ErrGraceShutdownTimeout = errors.New("server shutdown timed out")
)

type (
	// HTTPServer represents an HTTP server
	HTTPServer interface {
		Shutdown(ctx context.Context) error
		Serve(l net.Listener) error
	}

	// GRPCServer represents interface for grpc server
	GRPCServer interface {
		GracefulStop()
		Stop()
		Serve(l net.Listener) error
	}
)

// ServeHTTP start the HTTP server on the given address and add graceful shutdown handler.
//
// graceTimeout specify how long we want to wait for the shutdown to run.
// if graceTimeout = 0, we use default value: 30 seconds.
func ServeHTTP(srv HTTPServer, address string, graceTimeout time.Duration) error {
	// start graceful listener
	lis, err := Listen(address)
	if err != nil {
		return err
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = defaultGraceTimeout
		}

		stopped := make(chan struct{})
		ctx, cancel := context.WithTimeout(ctx, graceTimeout)
		defer cancel()

		go func() {
			srv.Shutdown(ctx)
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			return ErrGraceShutdownTimeout
		case <-stopped:
		}

		return nil
	})

	log.Printf("HTTP server is running on address: %s", address)

	// start serving
	if err := srv.Serve(lis); err != http.ErrServerClosed {
		return err
	}

	<-stoppedCh
	log.Println("HTTP server stopped")
	return nil

}

// ServeGRPC start the grpc server on the given address and add graceful shutdown handler.
//
// graceTimeout specify how long we want to wait for the graceful stop to run.
// if exceed the duration, we will forcefully stop the gRPC server.
// if graceTimeout = 0, we use default value: 30 seconds.
func ServeGRPC(server GRPCServer, address string, graceTimeout time.Duration) error {
	lis, err := Listen(address)
	if err != nil {
		return err
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = defaultGraceTimeout
		}

		stopped := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		select {
		case <-time.After(graceTimeout):
			server.Stop()
			return ErrGraceShutdownTimeout
		case <-stopped:
		}

		return nil
	})

	log.Printf("GRPC server is running on address %s", address)
	if err := server.Serve(lis); err != nil {
		return err
	}

	<-stoppedCh
	log.Println("GRPC server stopped")
	return nil
}

// WaitTermSig wait for termination signal and then execute the given handler when the signal received.
//
// The handler is usually service shutdown, so we can properly shutdown our server upon SIGTERM.
//
// It returns channel which will be closed after the signal received and the handler executed.
// We can use the signal to wait for the shutdown to be finished.
func WaitTermSig(handler func(context.Context) error) <-chan struct{} {
	stoppedCh := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// Wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, then shut down
		if err := handler(context.Background()); err != nil {
			log.Printf("graceful shutdown failed: %s", err.Error())
		} else {
			log.Println("graceful shutdown succeed")
		}

		close(stoppedCh)
	}()

	return stoppedCh
}

// Listen listens to the given port.
func Listen(port string) (net.Listener, error) {
	return net.Listen("tcp4", port)
}
