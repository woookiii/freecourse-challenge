package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type admin struct {
	network *Network
}

func (admin *admin) add(context *gin.Context) {
	res(context, http.StatusOK, "dshjkfdksdkjsfds")
}

func (admin *admin) update(context *gin.Context) {
}

func (admin *admin) viewAll(context *gin.Context) {
}

func (admin *admin) delete(context *gin.Context) {
}

func newAdmin(network *Network) {
	a := &admin{network: network}

	basePath := "/admin"

	network.register(basePath+"/add", POST, a.add)
	network.register(basePath+"/update", PUT, a.update)
	network.register(basePath+"/view", GET, a.viewAll)
	network.register(basePath+"/delete", DELETE, a.delete)
}
