package repository

import (
	"api/module/API/entity"
	"log"
	"time"

	"github.com/google/uuid"
)

func (r Repository) CreateMember(name string, email string, password string) error {
	if tx, err := r.db.Begin(); err != nil {
		return err
	} else if result, err := tx.Exec(
		"INSERT INTO member (id, name, email, password, role, created_time, updated_time, deleted_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		uuid.New().String(),
		name,
		email,
		password,
		"USER",
		time.Now(),
		nil,
		nil,
	); err != nil {
		tx.Rollback()
		log.Println("failed to insert member", err.Error())
		return err
	} else {
		count, _ := result.RowsAffected()
		tx.Commit()
		log.Println("Success to insert member", "count", count)
	}

	return nil
}

func (r Repository) FindMemberByEmail(email string) (*entity.Member, error) {
	var member entity.Member

	if err := r.db.QueryRow(
		"SELECT m.id, m.name, m.email, m.password, m.role, m.created_time, m.updated_time, m.deleted_time FROM member AS m WHERE m.email = ?",
		email,
	).Scan(
		&member.Id,
		&member.Name,
		&member.Email,
		&member.Password,
		&member.Role,
		&member.CreatedTime,
		&member.UpdatedTime,
		&member.DeletedTime,
	); err != nil {
		return nil, err
	} else {
		return &member, nil
	}
}
