// Package driver contains API implementations.
package driver

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Driver is an implementation of a server for our application.
type Driver interface {
	Serve(onExitCtx context.Context) error
}

// DriverKind encapsulates the application's different communication protocols.
type DriverKind int

const (
	Http DriverKind = iota
	Grpc
	GraphQL
)

var DriverNameToEnum = map[string]DriverKind{
	"http":    Http,
	"grpc":    Grpc,
	"graphql": GraphQL,
}

func httpGracefulServe(port int, onExitCtx context.Context, logger *zap.Logger) error {
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{Addr: addr}

	registerHttpGracefulShutdownHandler(server, onExitCtx, logger)

	logger.Info("start_listen", zap.Int("port", port))

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func registerHttpGracefulShutdownHandler(server *http.Server, onExitCtx context.Context, logger *zap.Logger) {
	go func() {
		<-onExitCtx.Done()
		logger.Warn("shutting_down_server")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("server_shutdown", zap.Error(err))
		} else {
			logger.Warn("server_exit")
		}
	}()
}
