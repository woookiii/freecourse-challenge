package connector

import (
	"log"
	"worker-search/config"
	"worker-search/connector/repository"
	"worker-search/connector/service"
	"worker-search/kafka"
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

	k := kafka.NewKafka(cfg, s)

	go func() {
		err := k.GetMessage([]string{"member"})
		if err != nil {
			log.Printf("Fail to get message : %v", err)
			return
		}
	}()

	return c
}
