package main

import (
	"net/http"

	"example.com/fxdemo/2-module/httpfx"
	"example.com/fxdemo/2-module/logfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		httpfx.Module,
		logfx.Module,

		fx.Provide(NewGraphPlotter),

		fx.Invoke(func(*http.Server) {}),
		fx.Invoke(func(g *GraphPlotter) {
			g.Plot()
		}),
	)

	app.Run()
}
