package network

import (
	"geosqlwithcdn/module/API/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	network *Network
}

func userRouter(network *Network) {
	u := &user{network}

	network.Router(POST, "/register-user", u.RegisterUser)
	network.Router(POST, "/upload-image", u.UploadImage)

	network.Router(GET, "/around-users", u.AroundUsers)

}

func (user *user) RegisterUser(context *gin.Context) {
	var req types.RegisterUserReq

	//ShouldBindJSON will get req body from context and bind it to our dto and check the binding did rightly
	if err := context.ShouldBindJSON(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, err.Error())
	} else if err = user.network.service.RegisterUser(req); err != nil {
		res(context, http.StatusInternalServerError, err.Error())
	} else {
		res(context, http.StatusOK, "Success")
	}
}

func (user *user) UploadImage(context *gin.Context) {
	name := context.Request.FormValue("username")
	file, header, err := context.Request.FormFile("image")

	if err != nil || name == "" {
		res(context, http.StatusUnprocessableEntity, err.Error())
	} else if err = user.network.service.UploadFile(name, header, file); err != nil {
		res(context, http.StatusInternalServerError, err.Error())
	} else {
		res(context, http.StatusOK, "Success")
	}
}

func (user *user) AroundUsers(context *gin.Context) {
	var req types.AroundUsersReq

	if err := context.ShouldBindQuery(&req); err != nil {
		res(context, http.StatusUnprocessableEntity, err.Error())
	} else if result, err := user.network.service.FindAroundUsers(req.UserName, req.Range, req.Limit); err != nil {
		res(context, http.StatusInternalServerError, err.Error())
	} else {
		res(context, http.StatusOK, result)
	}
}
