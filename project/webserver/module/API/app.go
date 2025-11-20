package API

import (
	"api/config"
	"api/kafka"
	"api/module/API/network"
	"api/module/API/repository"
	"api/module/API/service"
)

type API struct {
	config *config.Config
}

func NewAPI(
	cfg *config.Config,
	k *kafka.Kafka,
) *API {
	api := &API{cfg}

	r := repository.NewRepository(cfg)

	s := service.NewService(r, k)

	n := network.NewNetwork(cfg, s)

	go func() {
		n.Start()
		k.Close()
	}()

	return api
}
