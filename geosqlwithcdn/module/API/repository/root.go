package repository

import (
	"database/sql"
	"geosqlwithcdn/config"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	cfg *config.Config
	db  *sql.DB
}

func NewRepository(cfg *config.Config) *Repository {
	r := &Repository{cfg: cfg}
	var err error

	//connect db by passing driver name and path with port
	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		panic(err)
	} else if err = r.db.Ping(); err != nil {
		panic(err)
	} else {
		return r
	}

	return r
}
