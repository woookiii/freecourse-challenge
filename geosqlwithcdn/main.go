package main

import (
	"flag"
	"fmt"
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/db"
	"geosqlwithcdn/module/API"

	"go.uber.org/fx"
)

// go run . -cfg={flexible path}
var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	fx.New(
		//1. provide config by anonymous functions work for fx
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(func() *db.DBRoot { return db.RootDB(cfg) }),
		fx.Provide(func() *aws.Aws { return aws.NewAws(cfg) }),

		//2. provide target, can add more target if I need
		fx.Provide(API.NewAPI),

		//3. inject config dependencies
		fx.Invoke(func(_ *API.API) {}),
	).Run()
	config.NewConfig(*cfgPath)
	fmt.Println("start")
}
