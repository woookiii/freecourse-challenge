package API

import (
	"geosqlwithcdn/aws"
	"geosqlwithcdn/config"
	"geosqlwithcdn/module/API/network"
	"geosqlwithcdn/module/API/repository"
	"geosqlwithcdn/module/API/service"
)

type API struct {
	cfg *config.Config
	aws *aws.Aws
}

func NewAPI(
	cfg *config.Config,
	aws *aws.Aws,
) *API {
	api := &API{cfg, aws}

	r := repository.NewRepository(cfg)

	s := service.NewService(cfg, r, aws)

	n := network.NewNetwork(cfg, s)

	go func() {
		n.Start()
	}()

	return api
}
