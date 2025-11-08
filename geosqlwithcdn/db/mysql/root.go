package mysql

import (
	"database/sql"
	"geosqlwithcdn/config"
)

type DB struct {
	cfg *config.Config
	db  *sql.DB
}

func NewDB(cfg *config.Config) *DB {
	d := &DB{cfg: cfg}
	var err error

	//connect db by passing driver name and path with port
	if d.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		panic(err)
	} else if err = d.db.Ping(); err != nil {
		panic(err)
	} else {
		return d
	}

	return d
}
