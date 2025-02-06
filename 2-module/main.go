package main

import (
	"net/http"

	"github.com/gemini-kenshi/fx-demo/2-module/httpfx"
	"github.com/gemini-kenshi/fx-demo/graphfx"
	"github.com/gemini-kenshi/fx-demo/logfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		httpfx.Module,
		logfx.Module,
		fx.Provide(graphfx.NewGraphPlotter),

		fx.Invoke(func(*http.Server) {}),
		fx.Invoke(func(g *graphfx.GraphPlotter) {
			g.Plot()
		}),
	)

	app.Run()
}
