package connector

import (
	"worker/config"
	"worker/kafka"
	"worker/module/connector/repository"
	"worker/module/connector/service"
)

type Connector struct {
	config *config.Config
	kafka  *kafka.Kafka
}

func NewConnector(
	cfg *config.Config,
	k *kafka.Kafka,
) *Connector {
	c := &Connector{cfg, k}

	r := repository.NewRepository(cfg)

	service.NewService(r, k)

	return c
}
