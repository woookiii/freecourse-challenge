package repository

import (
	"context"
	"log"
	"worker-search/connector/document"
	"worker-search/connector/entity"
)

func (r *Repository) SaveMember(member *entity.Member) error {

	if result, err := r.typedClient.Exists("member", member.Id.String()).IsSuccess(nil); result {
		log.Println("This member document already exist")
		return nil
	} else if err != nil {
		log.Printf("Fail to check document existence")
		return err
	}

	d := document.Member{
		Id:          member.Id.String(),
		Name:        member.Name,
		Email:       member.Email,
		Password:    member.Password,
		Role:        member.Role,
		CreatedTime: member.CreatedTime,
		UpdatedTime: member.UpdatedTime,
		DeletedTime: member.DeletedTime,
	}

	res, err := r.typedClient.Index("member").
		Request(d).
		Do(context.Background())
	if err != nil {
		log.Printf("Fail to put document in member index, error: %v", err)
		return err
	}
	log.Printf("Success to put document in member index, result: %v", res.Result)
	return nil
}
