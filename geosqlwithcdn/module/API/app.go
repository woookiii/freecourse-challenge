package API

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/db"
	"geosqlwithcdn/module/API/network"
	"geosqlwithcdn/module/API/service"
)

type API struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

func NewAPI(
	cfg *config.Config,
	db *db.DBRoot,
	aws *aws.Aws,
) *API {
	api := &API{cfg, db, aws}

	s := service.NewService(cfg, db, aws)

	return api
}
