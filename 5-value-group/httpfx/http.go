package httpfx

import (
	"context"
	"net"
	"net/http"

	"github.com/gemini-kenshi/fx-demo/logfx"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(
		NewHTTPServer,

		fx.Annotate(
			NewServeMux,
			fx.ParamTags("", "", `group:"routes"`),
		),
		AsRoute(NewEchoHandler),
		AsRoute(NewHelloHandler),
	),
)

func NewHTTPServer(lc fx.Lifecycle, logger logfx.Logger, mux *http.ServeMux) *http.Server {
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

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(lc fx.Lifecycle, logger logfx.Logger, routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting ServeMux")
			return nil
		},
	})
	return mux
}
