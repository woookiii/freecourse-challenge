package authenticator

import (
	"encoding/base32"
	"net/url"
	"os"
	"qrauthscrapcron/config"

	"rsc.io/qr"
)

type authenticator struct {
	config       *config.Config
	secretBase32 string
}

type AuthenticatorImpl interface {
}

func NewAuthenticator(config *config.Config) (AuthenticatorImpl, error) {
	a := &authenticator{config: config}

	authCfg := config.Authenticator

	var secret []byte

	for _, char := range authCfg.Secret {
		secret = append(secret, byte(char))
	}

	a.secretBase32 = base32.StdEncoding.EncodeToString(secret)
	account := authCfg.Account
	issuer := authCfg.Issuer

	if URL, err := url.Parse("otpauth://totp"); err != nil {
		return nil, err
	} else {
		URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)

		params := url.Values{}
		params.Add("secret", a.secretBase32)
		params.Add("issuer", issuer)

		//TODO add param to URL

		if code, err := qr.Encode(URL.String(), qr.Q); err != nil {
			return nil, err
		} else if err = os.WriteFile(authCfg.FileName, code.PNG(), 0600); err != nil {
			return nil, err
		} else {
			return a, nil
		}
	}

}

func (a *authenticator) VerifySecret(secret string) (bool, error) {
	return false, nil
}
