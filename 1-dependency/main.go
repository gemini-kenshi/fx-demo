package main

import (
	"net/http"

	"github.com/gemini-kenshi/fx-demo/graphfx"
	"go.uber.org/fx"
)

func main() {
	// app := fx.New(
	// 	Module,
	// 	fx.Provide(NewLogger),
	// )

	app := fx.New(
		fx.Provide(
			NewHTTPServer,
			NewEchoHandler,
			NewServeMux,
		),
		fx.Provide(NewLogger),
		fx.Provide(graphfx.NewGraphPlotter),
		fx.Invoke(func(*http.Server) {}),
		fx.Invoke(func(g *graphfx.GraphPlotter) {
			g.Plot()
		}),
	)

	app.Run()
}
