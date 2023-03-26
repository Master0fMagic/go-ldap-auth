package ldap

import (
	"errors"
)

const (
	adminGroupName     = "admin-users"
	operatorGroupName  = "operator-users"
	userGroupAttribute = "ou"
)

var (
	ErrUserNotExists = errors.New("user not exists")
	ErrUserNotUnique = errors.New("user not exists")
)
