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

		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
