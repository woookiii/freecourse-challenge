package service

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/db"
)

type service struct {
	cfg *config.Config

	db  *db.DBRoot
	aws *aws.Aws
}

type ServiceImpl interface {
}

func NewService(
	cfg *config.Config,
	db *db.DBRoot,
	aws *aws.Aws,
) ServiceImpl {
	s := &service{cfg, db, aws}

	return s
}
