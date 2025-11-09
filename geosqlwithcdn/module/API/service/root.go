package service

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/db"
	"geosqlwithcdn/module/API/types"
	"log"
	"mime/multipart"
)

type service struct {
	config *config.Config

	db  *db.DBRoot
	aws *aws.Aws
}

type ServiceImpl interface {
	RegisterUser(req types.RegisterUserReq) error
	UploadFile(username string, header *multipart.FileHeader, file multipart.File) error
	FindAroundUsers(username string, searchRange int64, limit int64) (interface{}, error)
}

func NewService(
	config *config.Config,
	db *db.DBRoot,
	aws *aws.Aws,
) ServiceImpl {
	s := &service{config, db, aws}

	return s
}

func (service *service) RegisterUser(req types.RegisterUserReq) error {
	retryCount := 0

createAgain:
	if err := service.db.MySQL.RegisterUser(
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
		}
	}
	return nil
}

func (service *service) UploadFile(username string, header *multipart.FileHeader, file multipart.File) error {
	return nil
}

func (service *service) FindAroundUsers(username string, searchRange, limit int64) (interface{}, error) {
	return nil, nil
}
