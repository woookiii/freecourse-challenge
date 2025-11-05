package client

import (
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	client *grpc.ClientConn
	//we use AuthServiceClient is implemented by script
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	//dial(NewClient) is sort of connection to certain value, we can give option like NewCredentials
	// Establish a connection to the gRPC server using insecure credentials.
	if client, err := grpc.NewClient(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		// after making grpc client connection with url and security option
		// bind it to our client construct implemented by proto with script
		c.authClient = auth.NewAuthServiceClient(c.client)
		//this using polymorphism by ClientConnInterface

		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}

	return c, nil
}

//rpc CreateAuth(CreateTokenReq) returns (CreateTokenRes);
//rpc VerifyAuth(VerifyTokenReq) returns (VerifyTokenRes);

// we use type we define at proto
func (g *GRPCClient) CreateAuth(address string) (*auth.AuthData, error) {
	return nil, nil
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.VerifyTokenRes, error) {
	return nil, nil
}
