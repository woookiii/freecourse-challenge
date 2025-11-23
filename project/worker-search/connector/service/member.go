package service

import (
	"log"
	"worker-search/connector/entity"
)

func (s *Service) SaveMember(member *entity.Member) error {

	err := s.repository.SaveMember(member)
	if err != nil {
		log.Println("Failed to save member to replica db", "Member name", member.Name, "err", err)
		return err
	}

	log.Println("Success to save new member to replica db", "Member name", member.Name)

	return nil
}
