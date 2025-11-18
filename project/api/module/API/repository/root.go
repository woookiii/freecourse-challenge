package repository

import (
	"api/config"
	"api/module/API/types/entity"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Repository struct {
	config *config.Config
	db     *sql.DB
}

func (r Repository) CreateMember(name string, email string, password string) error {
	if tx, err := r.db.Begin(); err != nil {
		return err
	} else {
		id := uuid.New().String()
		if result, err := tx.Exec(
			"INSERT INTO member (id, name, email, password) VALUES (?,?,?,?);",
			id,
			name,
			email,
			password,
		); err != nil {
			tx.Rollback()
			return err
		} else {
			count, _ := result.RowsAffected()
			log.Println("Success to insert member", "count", count)
		}
		tx.Commit()
	}

	return nil
}

func (r Repository) FindMemberByEmail(email string) (*entity.Member, error) {
	var member entity.Member

	if err := r.db.QueryRow(
		"SELECT m.id, m.name, m.email, m.password FROM member AS m WHERE m.email = ?",
		email,
	).Scan(
		&member.Id,
		&member.Name,
		&member.Email,
		&member.Password,
	); err != nil {
		return nil, err
	}
	return &member, nil
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
