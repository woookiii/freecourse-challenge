package service

import (
	"encoding/json"
	"errors"
	"log"
	"worker-cache/connector/entity"
	"worker-cache/connector/repository"

	"github.com/IBM/sarama"
)

type Service struct {
	repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r}
}

func (s *Service) SaveMessage(message *sarama.ConsumerMessage) error {
	if message.Topic == "member" {
		var member entity.Member
		if err := json.Unmarshal(message.Value, &member); err != nil {
			log.Printf("Fail to unmarshal msg value to member struct: %v", err)
		}
		err := s.saveMember(&member)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("this topic not exist")
}
