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

func res(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, &response{
		header: &header{Result: s, Data: strings.Join(data, ",")},
		Result: res,
	})
}

func setGin(e *gin.Engine) {
	//see requests in console
	e.Use(gin.Logger())
	//if api dying by panic or something like that, restart it
	e.Use(gin.Recovery())
	e.Use(cors.New(cors.Config{
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

func (n *Network) Router(r Router, path string, handler gin.HandlerFunc) {
	e := n.e

	switch r {
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
