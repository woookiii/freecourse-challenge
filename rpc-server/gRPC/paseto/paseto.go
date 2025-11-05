package paseto

import (
	"rpc-server/config"
	auth "rpc-server/gRPC/proto"

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

func (m *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	//pointer receiver is another parameter, yet it gets pointer in front of method call

	//randomBytes := make([]byte, 16)
	//rand.Read(randomBytes)
	return m.Pt.Encrypt(m.Key, auth, nil)

	//auth is payload of token
}

func (m *PasetoMaker) VerifyToken(token string) error {

	var auth auth.AuthData
	m.Pt.Decrypt(token, m.Key, &auth, nil)
	//these two lines will check whether key is same and payload type is same

	return nil
}
