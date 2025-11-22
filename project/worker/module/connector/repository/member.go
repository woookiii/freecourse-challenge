package repository

import (
	"log"
	"worker/module/connector/entity"
)

func (r *Repository) FindMemberByEmail(email string) (*entity.Member, error) {
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

func (r *Repository) SaveMember(member *entity.Member) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec(
		"INSERT INTO member (id, name, email, password, role, created_time, updated_time, deleted_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		member.Id,
		member.Name,
		member.Email,
		member.Password,
		member.Role,
		member.CreatedTime,
		nil,
		nil,
	)
	if err != nil {
		tx.Rollback()
		log.Printf("failed to insert member: %v", err)
		return err
	}
	count, _ := result.RowsAffected()
	tx.Commit()
	log.Printf("Success to insert member - count: %v", count)

	return nil
}
