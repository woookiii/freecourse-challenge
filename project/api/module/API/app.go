package API

import (
	"api/config"
	"api/module/API/repository"
)

type API struct {
	config *config.Config
}

func NewAPI(
	cfg *config.Config,
) *API {
	api := &API{cfg}

	r := repository.NewRepository(cfg)

	return api
}
