package service

import (
	"worker/module/connector/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r}
}
