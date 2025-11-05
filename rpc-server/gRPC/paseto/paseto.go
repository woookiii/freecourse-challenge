package paseto

import (
	"rpc-server/config"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (m *PasetoMaker) CreateNewToken() (string, error) {
	//pointer receiver is another parameter, yet it gets pointer in front of method call
	return "", nil
}

func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
