package main

import (
	"flag"
	"worker-search/config"
	"worker-search/connector"

	"go.uber.org/fx"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	fx.New(
		fx.Provide(func() *config.Config { return cfg }),

		fx.Provide(connector.NewConnector),

		fx.Invoke(func(_ *connector.Connector) {}),
	).Run()
}
