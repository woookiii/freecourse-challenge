package main

import (
	"flag"
	"rpc-server/cmd"
	"rpc-server/config"
	"rpc-server/gRPC/server"
	"time"
)

// get flexible config
var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)

	server.NEWGRPCServer(cfg)

	time.Sleep(1e9)

	cmd.NewApp(cfg)
}
