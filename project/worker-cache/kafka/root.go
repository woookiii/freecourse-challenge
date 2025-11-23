package kafka

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"worker-cache/config"
	"worker-cache/connector/service"

	"github.com/IBM/sarama"
)

type Kafka struct {
	ready    chan bool
	consumer sarama.ConsumerGroup
	service  *service.Service
}

func NewKafka(cfg *config.Config, s *service.Service) *Kafka {
	consumer, err := connectConsumer(cfg, "redis-connector")
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	return &Kafka{
		ready:    make(chan bool),
		consumer: consumer,
		service:  s,
	}
}

func connectConsumer(config *config.Config, groupID string) (sarama.ConsumerGroup, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = config.Kafka.ClientId
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.Version = 1
	cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	cfg.Net.SASL.User = config.Kafka.APIKey
	cfg.Net.SASL.Password = config.Kafka.Secret
	cfg.Net.TLS.Enable = true
	cfg.Net.SASL.Handshake = true

	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	//if balance strategy need to be change flexible, use switch-case with config di
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	//this setting make possible to consume message which is stored but not consumed for certain reason like worker server down

	return sarama.NewConsumerGroup(config.Kafka.URLS, groupID, cfg)
}

func (k *Kafka) GetMessage(topics []string) error {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := k.consumer.Consume(ctx, topics, k); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Printf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
			k.ready = make(chan bool)
		}
	}()

	<-k.ready
	log.Println("Sarama consumer up and running")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true
	consumptionIsPaused := false
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(k.consumer, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()

	if err := k.consumer.Close(); err != nil {
		log.Printf("Error closing client: %v", err)
		return err
	}
	return nil
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
