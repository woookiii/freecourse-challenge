package service

import (
	"log"
	"qrauthscrapcron/repository"
	"qrauthscrapcron/types/schema"
)

type service struct {
	repository repository.RepositoryImpl
}

type ServiceImpl interface {
	Add(url, cardSelector, innerSelector string, tag []string) error
	View(url string) (*schema.Admin, error)
	ViewAll() ([]*schema.Admin, error)
	Update(url, cardSelector, innerSelector string, tag []string) error
	Delete(url string) error
}

func NewService(repository repository.RepositoryImpl) (ServiceImpl, error) {
	s := &service{repository: repository}

	return s, nil
}

func (s *service) Add(url, cardSelector, innerSelector string, tag []string) error {
	if err := s.repository.Add(url, cardSelector, innerSelector, tag); err != nil {
		log.Println("Failed to call add admin data", "err", err)
		return err
	}
	return nil
}

func (s *service) Update(url, cardSelector, innerSelector string, tag []string) error {
	if err := s.repository.Update(url, cardSelector, innerSelector, tag); err != nil {
		log.Println("Failed to call update admin data", "err", err)
	}
	return nil
}

func (s *service) ViewAll() ([]*schema.Admin, error) {
	if res, err := s.repository.ViewAll(); err != nil {
		log.Println("Failed to call view all admin data", "err", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (s *service) View(url string) (*schema.Admin, error) {
	if res, err := s.repository.View(url); err != nil {
		log.Println("Failed to call view admin data", "err", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (s *service) Delete(url string) error {
	if err := s.repository.Delete(url); err != nil {
		log.Println("Failed to call delete admin data", "err", err)
	}
	return nil
}
