package authenticator

import (
	"encoding/base32"
	"qrauthscrapcron/config"
)

type authenticator struct {
	config       *config.Config
	secretBase32 string
}

type AuthenticatorImpl interface {
}

func NewAuthenticator(config *config.Config) *AuthenticatorImpl {
	a := &authenticator{config: config}

	authCfg := config.Authenticator

	var secret []byte

	for _, char := range authCfg.Secret {
		secret = append(secret, byte(char))
	}

	a.secretBase32 = base32.StdEncoding.EncodeToString(secret)
	account := authCfg.Account
	issuer := authCfg.Issuer

	return a
}

func (a *authenticator) VerifySecret(secret string) (bool, error) {
	return false, nil
}
