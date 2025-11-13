package authenticator

import (
	"net/url"
	"os"
	"qrauthscrapcron/config"

	"github.com/dgryski/dgoogauth"
	"rsc.io/qr"
)

type authenticator struct {
	Secret   string
	Account  string
	Issuer   string
	FileName string
}

type AuthenticatorImpl interface {
	VerifySecret(secret string) (bool, error)
}

func NewAuthenticator(config *config.Config) (AuthenticatorImpl, error) {
	a := &authenticator{
		Secret:   config.Authenticator.Secret,
		Account:  config.Authenticator.Account,
		Issuer:   config.Authenticator.Issuer,
		FileName: config.Authenticator.FileName,
	}

	if URL, err := url.Parse("otpauth://totp"); err != nil {
		return nil, err
	} else {
		URL.Path += "/" + url.PathEscape(a.Account)

		params := url.Values{}
		params.Add("secret", a.Secret)
		params.Add("issuer", a.Issuer)
		URL.RawQuery = params.Encode()

		if code, err := qr.Encode(URL.String(), qr.Q); err != nil {
			return nil, err
		} else if err = os.WriteFile(a.FileName, code.PNG(), 0600); err != nil {
			return nil, err
		} else {
			return a, nil
		}
	}

}

func (a *authenticator) VerifySecret(password string) (bool, error) {
	otp := &dgoogauth.OTPConfig{
		Secret:     a.Secret,
		WindowSize: 1,
	}

	if valid, err := otp.Authenticate(password); err != nil {
		return false, err
	} else if !valid {
		return false, nil
	} else {
		return true, nil
	}
}
