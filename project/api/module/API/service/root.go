package service

import (
	"api/module/API/repository"
	"api/module/API/types/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *repository.Repository
}

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
		}
		log.Println("Success create new member", "Member name", req.Name)

		//TODO produce message to kafka in another goroutine

		return nil
	}

}

func NewService(r *repository.Repository) *Service {
	s := &Service{r}

	return s
}
