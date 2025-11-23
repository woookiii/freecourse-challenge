package service

import (
	"log"
	"worker-search/connector/entity"
)

func (s *Service) saveMember(member *entity.Member) error {

	err := s.repository.SaveMember(member)
	if err != nil {
		log.Printf("Failed to save member to elastic search Member id: %v, err: %v", member.Id, err)
		return err
	}

	log.Printf("Success to save new member to elastic search Member Id: %v", member.Id)

	return nil
}
