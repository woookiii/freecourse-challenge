package API

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/db"
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

	return api
}
