package connector

import (
	"log"
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
