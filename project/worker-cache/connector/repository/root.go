package repository

import "worker-cache/config"

type Repository struct {
}

func NewRepository(config *config.Config) *Repository {
	r := new(Repository)

	return r
}
