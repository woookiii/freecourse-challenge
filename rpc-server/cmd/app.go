package cmd

import (
	"rpc-server/config"
	"rpc-server/gRPC/client"
	"rpc-server/network"
	"rpc-server/repository"
	"rpc-server/service"
)

type App struct {
	cfg *config.Config

	gRPClient  *client.GRPCClient
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	// star(*) uses when touching pointer's value
	// and(&) uses passing the pointer of the declared and initiated instance
	a := &App{cfg: cfg}

	var err error

	client.NewGRPCClient(cfg)

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.gRPClient, err = client.NewGRPCClient(cfg); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPClient); err != nil {
		panic(err)
	} else {
		a.network.StartServer()
	}
}
