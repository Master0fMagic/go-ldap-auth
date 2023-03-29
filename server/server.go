package server

import (
	"github.com/Master0fMagic/go-ldap-auth/server/api"
	"github.com/gin-gonic/gin"
)

func Run(lh api.LoginHandler) {
	router := gin.Default()

	router.POST("/api/v1/login", lh.HandleLogin)

}
