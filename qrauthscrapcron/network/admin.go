package network

import (
	"net/http"
	"qrauthscrapcron/types"

	"github.com/gin-gonic/gin"
)

type admin struct {
	network *Network
}

func (admin *admin) add(context *gin.Context) {
	req := types.AddReq{}

	//ShouldBindJSON not only return error when binding is not correct, but also bind json to req
	if err := context.ShouldBindJSON(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, nil, err.Error())
	} else if err = admin.network.service.Add(req.URL, req.CardSelector, req.InnerSelector, req.Tag); err != nil {
		res(context, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(context, http.StatusOK, "Success", "Success")
	}

}

func (admin *admin) update(context *gin.Context) {
	req := types.UpdateReq{}

	if err := context.ShouldBindJSON(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, nil, err.Error())
	} else if err = admin.network.service.Add(req.URL, req.CardSelector, req.InnerSelector, req.Tag); err != nil {
		res(context, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(context, http.StatusOK, "Success", "Success")
	}

}

func (admin *admin) view(context *gin.Context) {
	req := types.ViewReq{}

	if err := context.ShouldBindJSON(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, nil, err.Error())
	} else if response, err := admin.network.service.View(req.URL); err != nil {
		res(context, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(context, http.StatusOK, response, "Success")
	}
}

func (admin *admin) viewAll(context *gin.Context) {

	if response, err := admin.network.service.ViewAll(); err != nil {
		res(context, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(context, http.StatusOK, response, "Success")
	}
}

func (admin *admin) delete(context *gin.Context) {
	req := types.DeleteReq{}

	if err := context.ShouldBindJSON(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, nil, err.Error())
	} else if err = admin.network.service.Delete(req.URL); err != nil {
		res(context, http.StatusInternalServerError, nil, err.Error())
	} else {
		res(context, http.StatusOK, "Success", "Success")
	}
}

func newAdmin(network *Network) {
	a := &admin{network: network}

	basePath := "/admin"

	network.register(basePath+"/add", POST, a.add)
	network.register(basePath+"/update", PUT, a.update)
	network.register(basePath+"/viewAll", GET, a.viewAll)
	network.register(basePath+"/view", GET, a.view)
	network.register(basePath+"/delete", DELETE, a.delete)
}
