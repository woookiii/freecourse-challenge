package service

import (
	"api/module/API/dto"
	"encoding/json"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateMember(req *dto.MemberSaveReq) error {
	if member, _ := s.repository.FindMemberByEmail(req.Email); member != nil {
		log.Println("This email already exist")
		return nil
	} else if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else if registeredMember, err := s.repository.CreateMember(
		req.Name,
		req.Email,
		string(hashedPassword),
	); err != nil {
		log.Println("Failed to create member", "Member name", req.Name, "err", err.Error())
		return err
	} else {
		log.Println("Success create new member", "Member name", req.Name)

		go func() {
			registeredMemberInBytes, err := json.Marshal(registeredMember)
			if err != nil {
				log.Println(err)
			}

			err = s.kafka.PushMessage("member", registeredMemberInBytes)
			if err != nil {
				log.Println(err)
			}
		}()

		return nil
	}

}
