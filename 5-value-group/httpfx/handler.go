package httpfx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gemini-kenshi/fx-demo/logfx"
	"go.uber.org/fx"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	logger logfx.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(lc fx.Lifecycle, logger logfx.Logger) *EchoHandler {
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

func (h *EchoHandler) Pattern() string {
	return "/echo"
}

// HelloHandler is an HTTP handler that
// prints a greeting to the user.
type HelloHandler struct {
	log logfx.Logger
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler(lc fx.Lifecycle, logger logfx.Logger) *HelloHandler {
	handler := &HelloHandler{log: logger}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HelloHandler")
			return nil
		},
	})
	return handler
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

// ServeHTTP handles an HTTP request to the /hello endpoint.
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Failed to read request: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Errorf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
