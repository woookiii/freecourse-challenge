package network

import (
	"api/config"
	"api/module/API/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	service *service.Service
	engine  *gin.Engine
	port    string
}

func NewNetwork(cfg *config.Config, s *service.Service) *Network {
	n := &Network{
		service: s,
		engine:  gin.New(),
		port:    cfg.Info.Port,
	}

	return n
}

func (network *Network) Start() error {
	return network.engine.Run(network.port)
}
