package network

import (
	"geosqlwithcdn/config"
	"geosqlwithcdn/module/API/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	service service.ServiceImpl
	engine  *gin.Engine
	port    string
	config  *config.Config
}

func NewNetwork(
	config *config.Config,
	service service.ServiceImpl,
) *Network {
	n := &Network{
		config:  config,
		service: service,
		engine:  gin.New(),
		port:    config.Info.Port,
	}

	return n
}

func (network *Network) Start() error {
	return network.engine.Run(network.port)
}
