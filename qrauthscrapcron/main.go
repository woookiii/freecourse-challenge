package main

import (
	"flag"
	"qrauthscrapcron/cmd/app"
	"qrauthscrapcron/config"
)

// go run . -config=<another-config-path>
var pathFlag = flag.String("config", "./config.toml", "set toml path")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)
}
