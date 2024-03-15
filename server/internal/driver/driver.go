// Package driver contains API implementations.
package driver

import (
	"context"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/server/internal/analytics"
	"go.uber.org/zap"
	"net/http"
	"strings"
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

func httpGracefulServe(port int, mux http.Handler, onExitCtx context.Context, logger *zap.Logger) error {
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{Addr: addr, Handler: mux}

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

// analyticsMiddleware wraps an HTTP handler to measure response times
func analyticsMiddleware(obs analytics.Collector, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path

		path = strings.ReplaceAll(path, "/", ":")
		metric := fmt.Sprintf("http_response_duration_%s_%s", method, path)

		defer analytics.MeasureDuration(obs, metric)()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
