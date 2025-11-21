package connector

import (
	"worker/config"
	"worker/kafka"
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

	return c
}
