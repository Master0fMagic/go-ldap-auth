package auth

import (
	"github.com/Master0fMagic/go-ldap-auth/dto"
	"github.com/Master0fMagic/go-ldap-auth/storage/ldap"
)

type Service struct {
	ldap *ldap.Client
}

func New(client *ldap.Client) *Service {
	return &Service{
		ldap: client,
	}
}

func (s *Service) AuthenticateUser(username, password string) (*dto.User, error) {
	err := s.ldap.Authenticate(username, password)
	//todo add check for invalid creds error
	if err != nil {
		return nil, ErrAuthenticateUser
	}

	userGroup, err := s.ldap.GetUserRole(nil, username)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Username: username,
		Role:     userGroup,
	}, nil
}
