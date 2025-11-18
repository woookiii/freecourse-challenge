package API

import "api/config"

type API struct {
	config *config.Config
}

func NewAPI(
	cfg *config.Config,
) *API {
	api := &API{cfg}

	return api
}
