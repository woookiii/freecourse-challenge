package network

import (
	"geosqlwithcdn/module/API/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	n *Network
}

func userRouter(n *Network) {
	u := &user{n}

	n.Router(POST, "/register-user", u.RegisterUser)
	n.Router(POST, "/upload-image", u.UploadImage)

	n.Router(GET, "/around-users", u.AroundUsers)

}

func (u *user) RegisterUser(c *gin.Context) {
	var req types.RegisterUserReq

	//ShouldBindJSON will get req body from context and bind it to our dto and check the binding did rightly
	if err := c.ShouldBindJSON(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	}
}

func (u *user) UploadImage(c *gin.Context) {
	name := c.Request.FormValue("userName")
	file, handler, err := c.Request.FormFile("image")

	if err != nil || name == "" {
		res(c, http.StatusUnprocessableEntity, err.Error())
	} else {
		//TODO via service, communicate with aws
	}
}

func (u *user) AroundUsers(c *gin.Context) {
	var req types.AroundUsersReq

	if err := c.ShouldBindQuery(&req); err != nil {
		res(c, http.StatusUnprocessableEntity, err.Error())
	}
}
