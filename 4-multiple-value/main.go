package main

import (
	"net/http"

	"example.com/fxdemo/4-multiple-value/httpfx"
	"example.com/fxdemo/graphfx"
	"example.com/fxdemo/logfx"
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
