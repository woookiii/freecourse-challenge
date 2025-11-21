package main

import (
	"api/config"
	"api/kafka"
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
		fx.Provide(func() *kafka.Kafka { return kafka.NewKafka(cfg) }),

		fx.Provide(API.NewAPI),

		fx.Invoke(func(_ *API.API) {}),
	).Run()
}
