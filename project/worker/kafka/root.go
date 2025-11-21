package kafka

import (
	"worker/config"

	"github.com/IBM/sarama"
)

type Kafka struct {
	consumer sarama.Consumer
}

func NewKafka(cfg *config.Config) *Kafka {
	consumer, err := connectConsumer(cfg.Kafka.URLS)
	if err != nil {
		panic(err)
	}
	return &Kafka{consumer: consumer}
}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = "consumer-client"
	cfg.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, cfg)
}
