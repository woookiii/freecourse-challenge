package service

import (
	"api/module/API/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) error {
	if member, _ := s.repository.FindMemberByEmail(req.Email); member != nil {
		log.Println("This email already exist")
		return nil
	} else if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		retryCount := 0
	createAgain:
		if s.repository.CreateMember(
			req.Name,
			req.Email,
			string(hashedPassword),
		); err != nil {
			retryCount++

			if retryCount < 3 {
				goto createAgain
			} else {
				log.Println("Failed to create member", "Member name", req.Name, "err", err.Error())
				return err
			}
		} else {
			log.Println("Success create new member", "Member name", req.Name)
			//TODO produce message to kafka in another goroutine
			return nil
		}

	}

}
