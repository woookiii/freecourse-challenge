package service

import (
	"api/kafka"
	"api/module/API/repository"
)

type Service struct {
	repository *repository.Repository
	kafka      *kafka.Kafka
}

func NewService(r *repository.Repository, k *kafka.Kafka) *Service {
	s := &Service{r, k}

	return s
}
