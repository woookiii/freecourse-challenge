package service

import (
	"log"
	"worker/module/connector/entity"
)

func (s *Service) saveMember(member *entity.Member) error {
	if member, _ := s.repository.FindMemberByEmail(member.Email); member != nil {
		log.Println("This email already exist")
		return nil
	}
	err := s.repository.SaveMember(member)
	if err != nil {
		log.Printf("Failed to save member to replica db member id: %v, error: %v", member.Id, err)
		return err
	}

	log.Printf("Success to save new member to replica db member id: %v", member.Id)

	return nil

}
