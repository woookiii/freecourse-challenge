package db

import (
	"geosqlwithcdn/config"
	"geosqlwithcdn/db/mysql"
)

type DBRoot struct {
	cfg   *config.Config
	MySQL *mysql.DB
}

func RootDB(cfg *config.Config) *DBRoot {
	root := &DBRoot{cfg: cfg}

	root.MySQL = mysql.NewDB(cfg)

	return root
}
