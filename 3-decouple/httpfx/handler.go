package httpfx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"example.com/fxdemo/logfx"
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
