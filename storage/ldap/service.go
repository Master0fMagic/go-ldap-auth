package ldap

import (
	"context"
	"fmt"
	"github.com/Master0fMagic/go-ldap-auth/common"
	"github.com/Master0fMagic/go-ldap-auth/config"
	"golang.org/x/exp/slices"
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

func (s *Service) GetUserRole(_ context.Context, username string) (common.Role, error) {
	filterStr := fmt.Sprintf(s.cfg.UserFiler, username)
	searchReq := ldap.NewSearchRequest(s.cfg.BaseDn, ldap.ScopeWholeSubtree, 0, 0, 0, false, filterStr, []string{}, nil)

	result, err := s.conn.Search(searchReq)
	if err != nil {
		return common.None, err
	}

	if len(result.Entries) == 0 {
		return common.None, ErrUserNotExists
	}
	if len(result.Entries) > 1 {
		return common.None, ErrUserNotUnique
	}
	// todo check password

	userEntry := result.Entries[0]
	return getUserGroup(userEntry.GetAttributeValues(userGroupAttribute)), nil
}

func getUserGroup(groups []string) common.Role {
	if slices.Contains(groups, adminGroupName) {
		return common.Admin
	}
	if slices.Contains(groups, operatorGroupName) {
		return common.Operator
	}
	return common.None
}

func (s *Service) Close() {
	s.conn.Close()
}
