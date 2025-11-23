package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func (k *Kafka) Setup(_ sarama.ConsumerGroupSession) error {
	close(k.ready)
	return nil
}

func (k *Kafka) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (k *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg := <-claim.Messages():
			session.MarkMessage(msg, "")
			err := k.service.SaveMessage(msg)
			if err != nil {
				log.Printf("Fail to save message: %v", err)
			}
			continue
		case <-session.Context().Done():
			return nil
		}
	}
}
