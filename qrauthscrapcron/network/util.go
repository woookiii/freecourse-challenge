package network

import (
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
	Result int    `json:"result"`
	Data   string `json:"data"`
}

type response struct {
	*header
	Result interface{} `json:"result"`
}

func res(context *gin.Context, statusCode int, result interface{}, data ...string) {
	context.JSON(statusCode, &response{
		header: &header{Result: statusCode, Data: strings.Join(data, ",")},
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
