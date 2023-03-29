package ldap

import (
	"context"
	"fmt"
	"github.com/Master0fMagic/go-ldap-auth/common"
	"github.com/Master0fMagic/go-ldap-auth/config"
	"golang.org/x/exp/slices"
	"gopkg.in/ldap.v2"
	"sync"
)

type Client struct {
	conn *ldap.Conn
	cfg  config.LdapConfig
	mtx  sync.Mutex
}

func New(config config.LdapConfig) (*Client, error) {
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

	return &Client{
		conn: conn,
		cfg:  config,
		mtx:  sync.Mutex{},
	}, err
}

func (s *Client) Authenticate(username, password string) (err error) {
	userDn := fmt.Sprintf(s.cfg.UserAuthPattern, username)

	s.mtx.Lock()
	defer s.mtx.Unlock()

	authConn, err := ldap.Dial("tcp", s.cfg.Url)
	if err != nil {
		return
	}
	defer authConn.Close()

	err = authConn.Bind(userDn, password)
	return
}

func (s *Client) GetUserRole(_ context.Context, username string) (common.Role, error) {
	filterStr := fmt.Sprintf(s.cfg.UserFilter, username)
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

func (s *Client) Close() {
	s.conn.Close()
}
