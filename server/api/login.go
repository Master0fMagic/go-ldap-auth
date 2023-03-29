package api

import (
	"github.com/Master0fMagic/go-ldap-auth/service/auth"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	authService auth.Service
}

func (h *LoginHandler) HandleLogin(ctx *gin.Context) {
	//get username and password from body
	//call auth service
	//return response
}
