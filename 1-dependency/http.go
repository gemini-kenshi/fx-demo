package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(
		NewHTTPServer,
		NewEchoHandler,
		NewServeMux,
	),
	fx.Invoke(func(*http.Server) {}),
)

func NewHTTPServer(lc fx.Lifecycle, logger Logger, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: "localhost:8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Info("Starting HTTP server at ", srv.Addr)
			// log.Printf("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	logger Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(lc fx.Lifecycle, logger Logger) *EchoHandler {
	handler := &EchoHandler{logger: logger}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting EchoHandler")
			return nil
		},
	})
	return handler
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("EchoHandler received a request: %v", r.Body)
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(lc fx.Lifecycle, logger Logger, echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting ServeMux")
			return nil
		},
	})
	return mux
}
