package main

import (
	"net/http"

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
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
