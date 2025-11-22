package service

import "worker-search/connector/repository"

type Service struct {
	repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r}
}
