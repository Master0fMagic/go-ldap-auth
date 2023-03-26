package ldap

import "context"

type IService interface {
	GetUser(ctx context.Context, username, password string)
}
