package main

import (
	"net/http"

	"example.com/fxdemo/3-decouple/httpfx"
	"example.com/fxdemo/logfx"
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
