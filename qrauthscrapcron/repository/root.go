package repository

import (
	"database/sql"
	"qrauthscrapcron/config"
	"qrauthscrapcron/types/schema"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

type RepositoryImpl interface {
	Add(url, cardSelector, innerSelector string, tag []string) error
	View(url string) (*schema.Admin, error)
	ViewAll() ([]*schema.Admin, error)
	Update(url, cardSelector, innerSelector string, tag []string) error
	Delete(url string) error
}

func NewRepository(cfg *config.Config) (RepositoryImpl, error) {
	r := new(repository)
	dbCfg := cfg.DB
	var err error
	if r.db, err = sql.Open(dbCfg.Database, dbCfg.URL); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (r *repository) Add(url, cardSelector, innerSelector string, tag []string) error {
	if _, err := r.db.Exec(
		"INSERT INTO Scrapping.Admin(url, tag, cardSelector, innerSelector) VALUES(?, ?, ? ,?)",
		url, strings.Join(tag, ""), cardSelector, innerSelector,
	); err != nil {
		return err
	}
	return nil
}

func (r *repository) View(url string) (*schema.Admin, error) {
	s := new(schema.Admin)
	err := r.db.QueryRow("SELECT * FROM Scrapping.Admin WHERE url = ?", url).Scan(
		&s.ID,
		&s.URL,
		&s.CardSelector,
		&s.InnerSelector,
		&s.Tag,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	return s, err
}

func (r *repository) ViewAll() ([]*schema.Admin, error) {
	if cursor, err := r.db.Query("SELECT * FROM Scrapping.Admin"); err != nil {
		return nil, err
	} else {
		defer cursor.Close()

		var result []*schema.Admin

		for cursor.Next() {
			s := new(schema.Admin)

			if err = cursor.Scan(
				&s.ID,
				&s.URL,
				&s.CardSelector,
				&s.InnerSelector,
				&s.Tag,
				&s.CreatedAt,
				&s.UpdatedAt,
			); err != nil {
				return nil, err
			} else {
				result = append(result, s)
			}
		}

		if len(result) == 0 {
			return []*schema.Admin{}, nil
		} else {
			return result, nil
		}
	}
}

func (r *repository) Update(url, cardSelector, innerSelector string, tag []string) error {
	if _, err := r.db.Exec("UPDATE Scrapping.Admin SET tag = ?, cardSelector = ?, innerSelector = ? WHERE url = ?",
		strings.Join(tag, " "), cardSelector, innerSelector, url,
	); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(url string) error {
	if _, err := r.db.Exec("DELETE FROM Scrapping.Admin WHERE url = ?", url); err != nil {
		return err
	}
	return nil
}
