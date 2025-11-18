package API

import (
	"api/config"
	"api/module/API/repository"
	"api/module/API/service"
)

type API struct {
	config *config.Config
}

func NewAPI(
	cfg *config.Config,
) *API {
	api := &API{cfg}

	r := repository.NewRepository(cfg)

	s := service.NewService(r)

	return api
}
