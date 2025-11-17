package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func unmarshalToFields(fields []interface{}, to ...interface{}) error {
	if len(fields) != len(to) {
		return errors.New("Field Length is not match")
	} else {
		for i, field := range fields {
			if err := json.Unmarshal(field.([]byte), to[i]); err != nil {
				return err
			}
		}

		return nil
	}
}
