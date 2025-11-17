package network

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type API_REQUEST uint8

const (
	GET API_REQUEST = iota
	POST
	PUT
	DELETE
)

type header struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type response struct {
	*header
	Result interface{} `json:"result"`
}

func res(context *gin.Context, statusCode int, result interface{}, message ...string) {
	context.JSON(statusCode, &response{
		header: &header{Status: statusCode, Message: strings.Join(message, ",")},
		Result: result,
	})
}

func (n *Network) register(path string, t API_REQUEST, h ...gin.HandlerFunc) gin.IRoutes {
	switch t {
	case GET:
		return n.engin.GET(path, h...)
	case POST:
		return n.engin.POST(path, h...)
	case PUT:
		return n.engin.PUT(path, h...)
	case DELETE:
		return n.engin.DELETE(path, h...)
	default:
		return nil
	}
}

func (n *Network) verifyAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if secret := getSecretByAuthorization(context); secret == "" {
			res(context, http.StatusUnauthorized, nil, "auth secret need")
		} else if valid, err := n.authenticator.VerifySecret(secret); err != nil {
			res(context, http.StatusUnauthorized, nil, err.Error())
			context.Abort()
		} else if !valid {
			res(context, http.StatusUnauthorized, nil, "not valid otp")
			context.Abort()
		} else {
			context.Next()
		}
	}
}

func getSecretByAuthorization(context *gin.Context) string {
	auth := context.Request.Header.Get("Authorization")
	authSliced := strings.Split(auth, " ")

	if len(authSliced) > 1 {
		return authSliced[1]
	} else {
		return ""
	}
}
