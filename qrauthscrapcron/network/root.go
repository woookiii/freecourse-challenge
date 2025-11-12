package network

import (
	"qrauthscrapcron/authenticator"
	"qrauthscrapcron/config"
	"qrauthscrapcron/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	config *config.Config
	engin  *gin.Engine

	service       service.ServiceImpl
	authenticator authenticator.AuthenticatorImpl
}

func NewNetwork(
	config *config.Config,
	service service.ServiceImpl,
	authenticator authenticator.AuthenticatorImpl,
) *Network {
	n := &Network{
		config:        config,
		service:       service,
		authenticator: authenticator,
		engin:         gin.New(),
	}

	newAdmin(n)
	
	return n
}

func (n *Network) Run() error {
	return n.engin.Run(n.config.Network.Port)
}
