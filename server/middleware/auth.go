package middleware

import (
	"github.com/Master0fMagic/go-ldap-auth/common"
	"github.com/Master0fMagic/go-ldap-auth/service/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(service auth.Service, roles []common.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//todo finish
	}
}
