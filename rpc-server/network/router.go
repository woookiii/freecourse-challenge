package network

import (
	"rpc-server/config"
	"rpc-server/gRPC/client"
	"rpc-server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg *config.Config

	service   *service.Service
	gRPClient *client.GRPCClient

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, gRPClient *client.GRPCClient) (*Network, error) {
	n := &Network{cfg: cfg, service: service, engin: gin.New(), gRPClient: gRPClient}

	return n, nil
}

func (n *Network) StartServer() {
	// if function start with capital, it means public, if func start with lower, it means private
	n.engin.Run(":9090")
}
