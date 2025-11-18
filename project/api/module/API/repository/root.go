package repository

import (
	"api/config"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	config *config.Config
	db     *sql.DB
}

func NewRepository(cfg *config.Config) *Repository {
	r := &Repository{config: cfg}

	var err error
	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		panic(err)
	} else if err = r.db.Ping(); err != nil {
		panic(err)
	}

	return r

}
