package ldap

import (
	"github.com/Master0fMagic/go-ldap-auth/config"
	"gopkg.in/ldap.v2"
)

type Service struct {
	conn *ldap.Conn
	cfg  config.LdapConfig
}

func New(config config.LdapConfig) (*Service, error) {
	conn, err := ldap.Dial("tcp", config.Url)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	if err = conn.Bind(config.BindUser, config.Password); err != nil {
		return nil, err
	}

	return &Service{
		conn: conn,
		cfg:  config,
	}, err
}

func (s *Service) Close() {
	s.conn.Close()
}
