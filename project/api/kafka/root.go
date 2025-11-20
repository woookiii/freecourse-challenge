package kafka

import (
	"api/config"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type Kafka struct {
	producer sarama.AsyncProducer
}

func NewKafka(cfg *config.Config) *Kafka {
	producer, err := connectProducer([]string{cfg.Kafka.URL})
	if err != nil {
		panic(err)
	}
	k := &Kafka{producer}
	return k
}

func (k *Kafka) PushMessage(topic string, message []byte) error {

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	k.producer.Input() <- &msg

	select {
	case succeedMsg := <-k.producer.Successes():
		log.Println("Success to produce message", succeedMsg)
		return nil
	case err := <-k.producer.Errors():
		log.Println("Failed to produce message:", err)
		return err
	}
}

func (k *Kafka) Close() error {
	return k.producer.Close()
}

func connectProducer(brokers []string) (sarama.AsyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = "member-client"
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Compression = sarama.CompressionZSTD
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Retry.Max = 30
	cfg.Producer.Retry.Backoff = time.Millisecond * 10
	//cfg.Producer.Idempotent = true
	//cfg.Producer.RequiredAcks = sarama.WaitForAll
	//cfg.Net.MaxOpenRequests = 1

	return sarama.NewAsyncProducer(brokers, cfg)
}
