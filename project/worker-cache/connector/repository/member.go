package repository

import (
	"context"
	"log"
	"time"
	"worker-cache/connector/entity"
)

func (r *Repository) SaveMember(member *entity.Member) error {
	ctx := context.Background()

	key := member.Id.String()
	resp := r.client.Do(ctx, r.client.B().Hset().Key(key).
		FieldValue().
		FieldValue("name", member.Name).
		FieldValue("email", member.Email).
		FieldValue("password", member.Password).
		FieldValue("role", member.Role).
		FieldValue("created_time", member.CreatedTime.Format(time.RFC3339Nano)).
		Build())
	if resp.Error() != nil {
		log.Printf("Failed to set member in redis : %v", resp.Error())
		return resp.Error()
	}
	return nil
}
