package kafka

import (
	"encoding/json"
	"log"
	"worker-cache/connector/entity"

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
			//TODO: call different methods for respective topics
			var member entity.Member
			if err := json.Unmarshal(msg.Value, &member); err != nil {
				log.Printf("Fail to unmarshal msg value to member struct: %v", err)
				continue
			}
			err := k.service.SaveMember(&member)
			if err != nil {
				log.Printf("Fail to save member to redis: %v", err)
				continue
			}
		case <-session.Context().Done():
			return nil
		}
	}
}
