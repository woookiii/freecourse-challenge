package network

import (
	"geosqlwithcdn/config"
	"geosqlwithcdn/module/API/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	s    service.ServiceImpl
	e    *gin.Engine
	port string
	cfg  *config.Config
}

func NewNetwork(
	cfg *config.Config,
	s service.ServiceImpl,
) *Network {
	n := &Network{
		cfg:  cfg,
		s:    s,
		e:    gin.New(),
		port: cfg.Info.Port,
	}

	return n
}

func (n *Network) Start() error {
	return n.e.Run(n.port)
}
