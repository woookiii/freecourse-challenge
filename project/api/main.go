package main

import (
	"api/config"
	"api/module/API"
	"flag"

	"go.uber.org/fx"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	fx.New(
		fx.Provide(func() *config.Config { return cfg }),

		fx.Provide(API.NewAPI),

		fx.Invoke(func(_ *API.API) {}),
	).Run()
}
