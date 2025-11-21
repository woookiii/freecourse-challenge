package repository

import (
	"database/sql"
	"worker/config"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) *Repository {
	r := new(Repository)

	var err error
	r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL)
	if err != nil {
		panic(err)
	}
	err = r.db.Ping()
	if err != nil {
		panic(err)
	}

	return r
}
