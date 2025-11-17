package network

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router int8

const (
	GET Router = iota
	POST
	DELETE
	PUT
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

func setGin(engine *gin.Engine) {
	//see requests in console
	engine.Use(gin.Logger())
	//if api dying by panic or something like that, restart it
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			//"PATCH",
		},
		//AllowHeaders: []string{
		//	"ORIGIN",
		//	"Content-Length",
		//	"Content-Type",
		//	"Access-Control-Allow-Headers",
		//	"Access-Control-Allow-Origin",
		//	"Authorization",
		//	"X-Requested-With",
		//	"expires",
		//},
		//ExposeHeaders: []string{
		//	"ORIGIN",
		//	"Content-Length",
		//	"Content-Type",
		//	"Access-Control-Allow-Headers",
		//	"Access-Control-Allow-Origin",
		//	"Authorization",
		//	"X-Requested-With",
		//	"expires",
		//},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
}

func (network *Network) Router(router Router, path string, handler gin.HandlerFunc) {
	e := network.engine

	//switch and case indent is same
	switch router {
	case GET:
		e.GET(path, handler)
	case POST:
		e.POST(path, handler)
	case PUT:
		e.PUT(path, handler)
	case DELETE:
		e.DELETE(path, handler)

	default:
		panic("Failed to register routers")
	}
}
