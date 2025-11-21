package connector

import (
	"worker/config"
	"worker/kafka"
	"worker/module/connector/repository"
	"worker/module/connector/service"
)

type Connector struct {
	config *config.Config
}

func NewConnector(
	cfg *config.Config,
) *Connector {
	c := &Connector{cfg}

	r := repository.NewRepository(cfg)

	s := service.NewService(r)
	
	k := kafka.Kafka{}

	go func() {
		s.Start()
	}()

	return c
}
