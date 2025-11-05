package client

import (
	"context"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
	"time"

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
func (g *GRPCClient) CreateAuth(req *auth.AuthData) (*auth.AuthData, error) {
	now := time.Now()
	expiredTime := now.Add(30 * time.Minute)

	a := &auth.AuthData{
		//when writing go, watch out to write local variable which might conflict with the public construct or import alias
		Name:       req.Name,
		CreateDate: now.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	if token, err := g.pasetoMaker.CreateNewToken(a); err != nil {
		return nil, err
	} else {
		a.Token = token

		//Background is empty context non-nil
		//CreateTokenReq is construct, not func, so we make construct CreateTokenReq and return its pointer
		if res, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: a}); err != nil {
			return nil, err
		} else {
			return res.Auth, nil
		}
	}
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {
	if res, err := g.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res.V, nil
	}
}
