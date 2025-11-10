package service

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/module/API/repository"
	. "geosqlwithcdn/module/API/repository/mysql/types"
	"geosqlwithcdn/module/API/types"
	"log"
	"mime/multipart"
)

type service struct {
	config     *config.Config
	repository *repository.Repository
	aws        *aws.Aws
}

type ServiceImpl interface {
	RegisterUser(req types.RegisterUserReq) error
	UploadFile(username string, header *multipart.FileHeader, file multipart.File) error
	FindAroundUsers(username string, searchRange int64, limit int64) ([]*User, error)
}

func NewService(
	config *config.Config,
	repository *repository.Repository,
	aws *aws.Aws,
) ServiceImpl {
	s := &service{
		config,
		repository,
		aws,
	}

	return s
}

func (service *service) RegisterUser(req types.RegisterUserReq) error {
	retryCount := 0

createAgain:
	if err := service.repository.RegisterUser(
		req.UserName,
		req.Description,
		req.Hobby,
		req.Latitude,
		req.Longitude,
	); err != nil {
		retryCount++

		if retryCount < 3 {
			goto createAgain
		} else {
			log.Println("Failed to create user", "user", req.UserName, "err", err.Error())
			return err
		}
	} else {
		log.Println("Success tp create new user", "name", req.UserName)
		return nil
	}
}

func (service *service) FindAroundUsers(userName string, searchRange, limit int64) ([]*User, error) {
	if limit == 0 {
		limit = 5
	}

	if u, err := service.getUser(userName); err != nil {
		return nil, err
	} else if users, err := service.repository.AroundUser(u.UserName, u.Latitude, u.Longitude, searchRange, limit); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (service *service) getUser(userName string) (*User, error) {

	if u, err := service.repository.GetUser(userName); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func (service *service) UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error {
	return nil
}
