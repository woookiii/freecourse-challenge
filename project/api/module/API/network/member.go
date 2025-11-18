package network

import (
	"api/module/API/types/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type member struct {
	network *Network
}

func memberRouter(network *Network) {
	m := &member{network}

	network.Router(POST, "/member/create", m.create)
}

func (member *member) create(ctx *gin.Context) {
	var req dto.MemberSaveReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res(ctx, http.StatusUnprocessableEntity, err.Error())
	} else if err = member.network.service.CreateMember(&req); err != nil {
		res(ctx, http.StatusInternalServerError, err.Error())
	}
	res(ctx, http.StatusOK, "Success create member")
}
