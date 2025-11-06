package network

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if t := getAuthToken(c); t == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			if _, err := n.gRPClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, nil)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")

	authSliced := strings.Split(authToken, " ")
	if len(authSliced) > 1 {
		token = authSliced[1]
	}

	return token
}
