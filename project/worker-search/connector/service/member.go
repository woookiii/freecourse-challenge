package service

import (
	"log"
	"worker-search/connector/entity"
)

func (s *Service) SaveMember(member *entity.Member) error {
	if member, _ := s.repository.FindMemberByEmail(member.Email); member != nil {
		log.Println("This email already exist")
		return nil
	}
	err := s.repository.SaveMember(member)
	if err != nil {
		log.Println("Failed to save member to replica db", "Member name", member.Name, "err", err)
		return err
	}

	log.Println("Success to save new member to replica db", "Member name", member.Name)

	return nil
}
