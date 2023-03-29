package dto

import "github.com/Master0fMagic/go-ldap-auth/common"

type User struct {
	Username string      `json:"username"`
	Role     common.Role `json:"Role"`
}
