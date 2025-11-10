package service

import (
	"errors"
	"fmt"
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/module/API/repository"
	. "geosqlwithcdn/module/API/repository/mysql/types"
	"geosqlwithcdn/module/API/types"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
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

	fileName := header.Filename
	fileExt := filepath.Ext(fileName)

	if !solveImageExtension(fileExt) {
		return errors.New("Failed to solve extension")
	} else {
		path := "./upload-image"
		filePath := fmt.Sprintf("%s/%s", path, fileName)

		//create return readable and also writable file(which is new or truncated)
		if out, err := os.Create(filePath); err != nil {
			return err
		} else {
			defer out.Close()

			if _, err := io.Copy(out, file); err != nil {
				return err
			} else if service.putFileToS3(
				fileName,
				userName,
				strings.TrimPrefix(fileExt, "."),
				filePath,
			); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}

func (s *service) putFileToS3(fileName, userName, extension, path string) error {
	key := userName + "/" + fileName

	//open return read-only file
	if f, err := os.Open(path); err != nil {
		return err
	} else {
		defer f.Close()

		if err = s.aws.PutFileToS3(key, extension, f); err != nil {
			return err
		} else {
			return nil
		}
	}
}
