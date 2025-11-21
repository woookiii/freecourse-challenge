package service

import (
	"worker/kafka"
	"worker/module/connector/repository"
)

type Service struct {
	repository *repository.Repository
	kafka      *kafka.Kafka
}

func NewService(r *repository.Repository, k *kafka.Kafka) *Service {
	return &Service{r, k}
}
