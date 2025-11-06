package server

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
	"rpc-server/config"
	"rpc-server/gRPC/paseto"
	auth "rpc-server/gRPC/proto"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NEWGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})
		//register server we will use
		reflection.Register(server)

		go func() {
			log.Println("Start grpc server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()
		//if Serve is run, the codes after it can't run in same thread, so we Serve server in new thread through go routine
	}

	return nil
}

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token

	s.tokenVerifyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token

	res := &auth.VerifyTokenRes{V: &auth.Verify{
		Auth: nil,
	}}

	if authData, ok := s.tokenVerifyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED
		return res, errors.New("Token not exist")
	} else if err := s.pasetoMaker.VerifyToken(token); err != nil {
		return nil, errors.New("Invalid token value")
	} else if authData.ExpireDate < time.Now().Unix() {
		delete(s.tokenVerifyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE
		return res, errors.New("Expired time over")
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
		return res, nil
	}

}
